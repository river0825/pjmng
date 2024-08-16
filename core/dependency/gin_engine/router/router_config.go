package router

import (
	"github.com/gin-gonic/gin"

	gorm "river0825/cleanarchitecture/core/dependency/storage/coregorm"

	"river0825/cleanarchitecture/core/arch/domain/eventbus"
	"river0825/cleanarchitecture/core/arch/port/gin/middleware"
	"river0825/cleanarchitecture/core/dependency/settings"
)

type RouterGroup struct {
	ApiGroup       gin.IRouter
	PublicApiGroup gin.IRouter
	CmsGroup       gin.IRouter
	devGroup       gin.IRouter
	env            string
}

func NewRouterGroup(
	env string,
	ApiGroup gin.IRouter,
	PublicApiGroup gin.IRouter,
	CmsGroup gin.IRouter,
	DevGroup gin.IRouter,
) RouterGroup {
	return RouterGroup{
		ApiGroup:       ApiGroup,
		PublicApiGroup: PublicApiGroup,
		CmsGroup:       CmsGroup,
		devGroup:       DevGroup,
		env:            env,
	}
}

func (r RouterGroup) AddDevRoute(method string, path string, handler gin.HandlerFunc) {
	if r.env != "production" {
		r.devGroup.Handle(method, path, handler)
	}
}

type IRouter interface {
	RegisterRoutes(routerGroup RouterGroup)
}

type Dependencies struct {
	Config                  *settings.Config
	AuthorizationMiddleware *middleware.AuthorizationMiddleware
	EventBus                eventbus.IEventBus
	GormDB                  *gorm.DB
}
