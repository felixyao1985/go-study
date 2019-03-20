package index

import (
	"study/go-study/router"
)

var Routes = []router.Route{
	{
		"index",
		"GET",
		"/index",
		Index,
	},
	{
		"hello",
		"GET",
		"/hello/:name",
		Hello,
	},
}
