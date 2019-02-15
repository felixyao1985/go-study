package index

import (
	"go-study/router"
)

var Routes = []router.Route{
	router.Route{
		"index",
		"GET",
		"/index",
		Index,
	},
	router.Route{
		"hello",
		"GET",
		"/hello/:name",
		Hello,
	},
}
