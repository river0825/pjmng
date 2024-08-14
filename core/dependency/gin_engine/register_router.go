package gin_engine

import (
	"river0825/cleanarchitecture/core/dependency/gin_engine/router"
)

type RegisterRouteFunc func(dependency *router.Dependencies) router.IRouter

var regFunc []RegisterRouteFunc

func RegisterRoutes(f RegisterRouteFunc) {
	regFunc = append(regFunc, f)
}

func registerRoutes(dependency *router.Dependencies) []router.IRouter {
	var routes []router.IRouter
	for _, r := range regFunc {
		routes = append(routes, r(dependency))
	}
	return routes
}
