
package app

import ("go-study/app/home"
"go-study/app/index"

	"go-study/router"
)



func mergeRoutes()  []router.Route{
	routes := []router.Route{}
routes = merge(routes, home.Routes, index.Routes)
		return routes
	}