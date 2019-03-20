package home

import (
	"study/go-study/router"
)

var Routes = []router.Route{
	{
		"home_index",
		"GET",
		"/home_index",
		Index,
	},
	{
		"home_hello",
		"GET",
		"/home_hello/:name",
		Hello,
	},
}
