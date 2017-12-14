package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetAllMember",
		"GET",
		"/member",
		GetAllMember,
	},
	Route{
		"GetAllBook",
		"GET",
		"/book",
		GetAllBook,
	},
	Route{
		"GetMember",
		"GET",
		"/member/{member_id}",
		GetMember,
	},
	Route{
		"GetBook",
		"GET",
		"/book/{isbn}",
		GetBook,
	},
	Route{
		"GetMemberShelf",
		"GET",
		"/member/{member_id}/shelf",
		GetMemberShelf,
	},
	Route{
		"GetBookOwnedBy",
		"GET",
		"/book/{isbn}/owned-by",
		GetBookOwnedBy,
	},
	Route {
		"AddMember",
		"POST",
		"/member/add",
		AddMember,
	},
	Route {
		"AddBook",
		"POST",
		"/book/add",
		AddBook,
	},
	Route {
		"AddToShelf",
		"POST",
		"/member/{member_id}/shelf/add",
		AddToShelf,
	},
}