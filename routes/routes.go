package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// NewRouter prepare all routes to server
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	for _, route := range routeList {
		r.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return r
}

type Routes []Route

var routeList = Routes{
	Route{
		Name:        "getCustomers",
		Method:      "GET",
		Pattern:     "/customers",
		HandlerFunc: getCustomers,
	},
	Route{
		Name:        "getCustomer",
		Method:      "GET",
		Pattern:     "/customers/{id}",
		HandlerFunc: getCustomer,
	},
	Route{
		Name:        "addCustomer",
		Method:      "POST",
		Pattern:     "/customers",
		HandlerFunc: addCustomer,
	},
	Route{
		Name:        "updateCustomer",
		Method:      "PATCH",
		Pattern:     "/customers/{id}",
		HandlerFunc: updateCustomer,
	},
	Route{
		Name:        "deleteCustomer",
		Method:      "DELETE",
		Pattern:     "/customers/{id}",
		HandlerFunc: deleteCustomer,
	},
	Route{
		Name:        "index",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: index,
	},
}
