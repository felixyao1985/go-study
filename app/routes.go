
package app

import ("restfulApi/app/home"
"restfulApi/app/index"

	"restfulApi/router"
)



func mergeRoutes()  []router.Route{
	routes := []router.Route{}
routes = merge(routes, home.Routes, index.Routes)
		return routes
	}