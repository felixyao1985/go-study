package home

import (
	"go-study/router"
)

var Routes = []router.Route{
	router.Route{
		"home_index",
		"GET",
		"/home_index",
		Index,
	},
	router.Route{
		"home_hello",
		"GET",
		"/home_hello/:name",
		Hello,
	},
}
