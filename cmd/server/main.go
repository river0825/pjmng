package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/alecthomas/kong"
	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"river0825/cleanarchitecture/core/dependency/eventbus/inmem"
	"river0825/cleanarchitecture/core/dependency/gin_engine"
	"river0825/cleanarchitecture/core/dependency/gin_engine/checkalive"
	traceLog "river0825/cleanarchitecture/core/dependency/log"
	"river0825/cleanarchitecture/core/dependency/settings"
	"river0825/cleanarchitecture/core/dependency/storage/gorm"
)

const (
	defaultJobNum = 1
)

// The command-line arguments and flags, parsed by the Go kong package.
type CLI struct {
	// the version flag
	Version kong.VersionFlag `short:"V" name:"version" help:"Show version information and exit."`

	// the customized flags
	Debug  bool   `name:"debug" help:"Enable debug mode (syntax sugar)."`
	Config string `short:"c" name:"config" type:"path" default:"config.local.yaml" help:"The config file path."`
	Action string `short:"a" name:"action" default:"serve" enum:"create-index,serve,gen-token" help:"The perform action."`
}

func (c *CLI) ParseAndExit(ctx context.Context, wg *sync.WaitGroup, cancelFunc context.CancelFunc) {
	options := []kong.Option{
		kong.Description("The Nikeyi Server"),
		kong.Vars{
			"version": "0.0.1",
		},
	}

	// parse the command-line arguments and options
	kong.Parse(c, options...)

	// run the program by the parsed arguments and options
	if err := c.Run(ctx, wg, cancelFunc); err != nil {
		// Fatal will call os.Exit(1) after logging
		log.Fatal().Err(err).Msg("cannot run the command")
	}
}

func (c *CLI) Run(ctx context.Context, wg *sync.WaitGroup, cancelFunc context.CancelFunc) error {
	defer cancelFunc()
	defer wg.Done()
	// init config
	config, err := settings.NewConfig(c.Config)
	if err != nil {
		log.Error().Err(err).Msg("init config error")
		return err
	}

	// setup everything
	c.setupLogger(config)
	c.setupDebugMode()
	c.setupSentry(config)

	return c.run(ctx, config)
}

func (c *CLI) run(ctx context.Context, config *settings.Config) error {
	defer c.cleanup()

	switch c.Action {
	// case "gen-token":
	// 	authService := service.NewAuthService(config.Auth.JWTSecret)
	// 	token := authService.CreateAccessToken(ctx, "1")
	// 	fmt.Println(token)
	// 	return nil
	case "serve":
		// gorm
		// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		// 	config.DB.Host, config.DB.Port, config.DB.User, config.DB.Password, config.DB.DBName, config.DB.Schema)
		var db *gorm.DB
		if config.Logger.Sql {
			db = gorm.NewDBWithLogger("")
		} else {
			db = gorm.NewDB("")
		}

		// log.Info().Msgf("database connected, host %s, port %d, user %s, dbname %s", config.DB.Host, config.DB.Port, config.DB.User, config.DB.DBName)

		// set up auth api proxy
		// authService := service.NewAuthService(config.Auth.JWTSecret)

		eventBus := inmem.NewInMemEventBus()

		checkAlive := []checkalive.ICheckAlive{
			db,
		}

		// set up gin engine
		ginEngine := gin_engine.NewGinEngine(
			config,
			nil,
			eventBus,
			db,
			checkAlive,
		)
		return ginEngine.Serve(ctx)
	default:
		return fmt.Errorf("invalid action: %s", c.Action)
	}
}

func (c *CLI) setupLogger(config *settings.Config) {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logLevels := map[string]zerolog.Level{
		"panic": zerolog.PanicLevel,
		"fatal": zerolog.FatalLevel,
		"error": zerolog.ErrorLevel,
		"warn":  zerolog.WarnLevel,
		"info":  zerolog.InfoLevel,
		"debug": zerolog.DebugLevel,
		"trace": zerolog.TraceLevel,
	}
	level, exists := logLevels[config.Logger.Level]
	if !exists {
		// set the default log level to info
		level = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(level)

	if config.Env == "local" {
		zerolog.ErrorStackMarshaler = traceLog.DebugMarshalStack
	} else {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	}

	if config.Logger.Debug {
		writer := zerolog.ConsoleWriter{Out: os.Stderr}
		log.Logger = zerolog.New(writer).With().Timestamp().Logger()
	}
}

func (c *CLI) setupDebugMode() {
	if c.Debug {
		// enable the log level to trace and show pretty logger
		zerolog.SetGlobalLevel(zerolog.TraceLevel)

		writer := zerolog.ConsoleWriter{Out: os.Stderr}
		log.Logger = zerolog.New(writer).With().Timestamp().Logger()

		log.Debug().Msg("debug mode enabled")
	}
}

func (c *CLI) setupSentry(config *settings.Config) {
	if config.Logger.Sentry.DSN == "" {
		log.Info().Msg("sentry is not configured, skip")
		return
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: config.Logger.Sentry.DSN,
		// setup the tracing and sampling rate
		EnableTracing:      config.Logger.Sentry.Sampling > 0.0,
		TracesSampleRate:   config.Logger.Sentry.Sampling,
		ProfilesSampleRate: config.Logger.Sentry.Sampling,

		// Enable printing of SDK debug messages.
		Debug:            true,
		AttachStacktrace: true,
	})

	if err != nil {
		log.Error().Err(err).Msg("init sentry error, but keep server running")
		return
	}

	log.Debug().Msg("sentry is configured")
}

func (c *CLI) cleanup() {
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	sentry.Flush(2 * time.Second)
}

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	cli := CLI{}
	wg := &sync.WaitGroup{}
	wg.Add(defaultJobNum)
	go cli.ParseAndExit(ctx, wg, cancel)
	select {
	case <-quit:
		cancel()
		log.Info().Msg("server starting gracefully shutting down")
	case <-ctx.Done():
	}
	wg.Wait()
}
