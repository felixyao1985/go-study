package app

import (
	"study/go-study/app/home"
	"study/go-study/app/index"

	"study/go-study/router"
)

func mergeRoutes() []router.Route {
	routes := []router.Route{}
	routes = merge(routes, home.Routes, index.Routes)
	return routes
}
