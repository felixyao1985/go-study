
package app

import ("../app/home"
"../app/index"

	"../router"
)



func mergeRoutes()  []router.Route{
	routes := []router.Route{}
routes = merge(routes, home.Routes, index.Routes)
		return routes
	}