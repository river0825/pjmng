package gin_engine

import (
	"context"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"river0825/cleanarchitecture/core/arch/domain/eventbus"
	"river0825/cleanarchitecture/core/arch/port/gin/middleware"
	"river0825/cleanarchitecture/core/arch/port/gin/service"
	"river0825/cleanarchitecture/core/dependency/gin_engine/checkalive"
	"river0825/cleanarchitecture/core/dependency/gin_engine/router"
	"river0825/cleanarchitecture/core/dependency/settings"
	"river0825/cleanarchitecture/core/dependency/storage/coregorm"
)

var (
	// skip path list for the gin access logger
	skipLogPaths = []string{
		"/v2/livez",
		"/v2/readyz",
	}
)

type GinEngine struct {
	engine       *gin.Engine
	dependencies *router.Dependencies
	checkAlive   []checkalive.ICheckAlive
}

func NewGinEngine(
	config *settings.Config,
	authService service.JWTTokenAuthor,
	eventBus eventbus.IEventBus,
	gormDB *coregorm.DB,
	checkAlive []checkalive.ICheckAlive,
) *GinEngine {
	// create the customized gin engine
	engine := gin.New()
	engine.Use(
		// log the request, and skip the health check endpoint
		gin.LoggerWithWriter(gin.DefaultWriter, skipLogPaths...),
		// recover from any panics and return 500
		gin.Recovery(),
	)

	return &GinEngine{
		engine:     engine,
		checkAlive: checkAlive,
		dependencies: &router.Dependencies{
			Config:   config,
			EventBus: eventBus,
			GormDB:   gormDB,
			AuthorizationMiddleware: middleware.NewAuthorizationMiddleware(
				authService,
			),
		},
	}
}

func (e *GinEngine) setMiddleware() {
	// setup the gin/zerolog middleware
	e.engine.Use(logger.SetLogger(
		logger.WithSkipPath(skipLogPaths),
	))

	// setup the sentry middleware
	e.engine.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	e.engine.Use(e.CorsOptions)
}

func (e *GinEngine) CorsOptions(c *gin.Context) {
	// Split the allowed origins string into a slice
	allowedOriginsStr := e.dependencies.Config.App.AccessControlAllowOrigin
	allowedOriginsArr := strings.Split(allowedOriginsStr, ",")

	// Create a map to store the allowed origins for quick lookup
	allowedOrigins := make(map[string]bool)
	for _, origin := range allowedOriginsArr {
		allowedOrigins[strings.TrimSpace(origin)] = true
	}

	// Get the origin of the incoming request
	origin := c.Request.Header.Get("Origin")
	if allowedOrigins[origin] {
		c.Header("Access-Control-Allow-Origin", origin)
	}

	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}

func (e *GinEngine) ServeHTTP(w http.ResponseWriter, req *http.Request) error {
	e.setMiddleware()
	// set up routes here
	e.registerBasicEndpoint()
	e.registerRouters(registerRoutes(e.dependencies))

	e.engine.ServeHTTP(w, req)

	return nil
}

func (e *GinEngine) Serve(ctx context.Context) error {
	e.setMiddleware()
	// set up routes here
	e.registerBasicEndpoint()
	e.registerRouters(registerRoutes(e.dependencies))

	srv := &http.Server{
		Addr:    e.dependencies.Config.App.Port,
		Handler: e.engine,
	}

	// start server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("failed to start server")
			os.Exit(1)
		}
	}()

	log.Info().Str("port", e.dependencies.Config.App.Port).Msg("server started")

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		_ = srv.Close()
		return err
	}

	log.Info().Msg("server gracefully stopped")
	return nil
}

func (e *GinEngine) registerRouters(routers []router.IRouter) {
	if len(routers) == 0 {
		log.Warn().Msgf("[CustomRouter] custom module is empty")
		return
	}

	v2 := e.engine.Group("/v2")

	dev := v2.Group("/dev")
	dev.Use(e.dependencies.AuthorizationMiddleware.DevBindUserInfo())

	routerGroup := router.NewRouterGroup(
		e.dependencies.Config.Env,
		v2.Group("/api"),
		v2.Group("/pub"),
		v2.Group("/cms"),
		dev,
	)

	// setup auth access token middleware for endpoint under api group
	authMiddleware := e.dependencies.AuthorizationMiddleware
	routerGroup.ApiGroup.Use(authMiddleware.ValidateAccessToken())
	routerGroup.CmsGroup.
		Use(authMiddleware.ValidateAccessToken()).
		Use(authMiddleware.CmsAllowUsers())

	for _, r := range routers {
		r.RegisterRoutes(routerGroup)
		log.Info().Msgf("[CustomRouter] register module %s successfully", reflect.TypeOf(r))
	}
}

// register the basic API endpoint, like /readyz and /livez
// ref: https://kubernetes.io/docs/reference/using-api/health-checks/
func (e *GinEngine) registerBasicEndpoint() {
	// /livez     the basic API endpoint and always return ok
	e.engine.GET("/v2/livez", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// /readyz      the health check API endpoint
	e.engine.GET("/v2/readyz", func(c *gin.Context) {
		// check the database connection
		for _, srv := range e.checkAlive {
			if err := srv.Ping(c.Request.Context()); err != nil {
				log.Warn().Err(err).Msgf("%s connection error", srv.Name())
				c.String(http.StatusInternalServerError, "%s connection error")
				return
			}
		}
		c.String(http.StatusOK, "ok")
	})
}
