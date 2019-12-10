package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var control = &Controller{}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"Create Element in Cache",
		"POST",
		"/add",
		control.save,
	},
	Route{
		"Get Element from Cache",
		"GET",
		"/fetch/{key}",
		control.get,
	},
	Route{
		"Get All Element from Cache",
		"GET",
		"/fetchall",
		control.getAll,
	},
	Route{
		"Update Element from Cache",
		"PUT",
		"/update",
		control.update,
	},
	Route{
		"Remove Element from Cache",
		"DELETE",
		"/remove/{key}",
		control.remove,
	},
	/*Route{
		"Greeting",
		"GET",
		"/greet/{name}",
		control.greet,
	}*/
	}

	func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
	var handler http.Handler
	log.Println(route.Name)
	handler = route.HandlerFunc

	router.
	Methods(route.Method).
	Path(route.Pattern).
	Name(route.Name).
	Handler(handler)
	}
	return router
	}
