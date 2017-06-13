package src

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		BucketsIndex,
	},
	Route{
		"BucketIndex",
		"GET",
		"/{key}",
		BucketIndex,
	},
	Route{
		"BucketIndex",
		"POST",
		"/{key}",
		BucketCreate,
	},
	Route{
		"BucketDelete",
		"DELETE",
		"/{key}",
		BucketDelete,
	},
	Route{
		"ItemShow",
		"GET",
		"/{key}/{itemKey}",
		ItemShow,
	},
	Route{
		"ItemSet",
		"PUT",
		"/{key}/{itemKey}",
		ItemSet,
	},
	Route{
		"ItemCreate",
		"POST",
		"/{key}/{itemKey}",
		ItemSet,
	},
	Route{
		"ItemDelete",
		"DELETE",
		"/{key}/{itemKey}",
		ItemDelete,
	},
	/*  ,
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},*/
}
