/*
 * Appointment Scheduling Service- OpenAPI 3.0
 *
 * The specification contains API contractrs for appointment scheduling service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
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
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"CreateAppointment",
		strings.ToUpper("Post"),
		"/v1/appointment",
		CreateAppointment,
	},

	Route{
		"DeleteAppointmentById",
		strings.ToUpper("Delete"),
		"/v1/appointment/{id}",
		DeleteAppointmentById,
	},

	Route{
		"GetAppointmentById",
		strings.ToUpper("Get"),
		"/v1/appointment/{id}",
		GetAppointmentById,
	},

	Route{
		"GetAppointments",
		strings.ToUpper("Get"),
		"/v1/appointment",
		GetAppointments,
	},

	Route{
		"UpdateAppointmentById",
		strings.ToUpper("Put"),
		"/v1/appointment/{id}",
		UpdateAppointmentById,
	},

	Route{
		"CreateCustomer",
		strings.ToUpper("Post"),
		"/v1/customer",
		CreateCustomer,
	},

	Route{
		"DeleteCustomerById",
		strings.ToUpper("Delete"),
		"/v1/customer/{id}",
		DeleteCustomerById,
	},

	Route{
		"GetCustomerById",
		strings.ToUpper("Get"),
		"/v1/customer/{id}",
		GetCustomerById,
	},

	Route{
		"GetCustomers",
		strings.ToUpper("Get"),
		"/v1/customer",
		GetCustomers,
	},

	Route{
		"UpdateCustomerById",
		strings.ToUpper("Put"),
		"/v1/customer/{id}",
		UpdateCustomerById,
	},

	Route{
		"CreateItem",
		strings.ToUpper("Post"),
		"/v1/Item",
		CreateItem,
	},

	Route{
		"DeleteItemById",
		strings.ToUpper("Delete"),
		"/v1/item/{id}",
		DeleteItemById,
	},

	Route{
		"GetItemById",
		strings.ToUpper("Get"),
		"/v1/item/{id}",
		GetItemById,
	},

	Route{
		"GetItems",
		strings.ToUpper("Get"),
		"/v1/Item",
		GetItems,
	},

	Route{
		"UpdateItemById",
		strings.ToUpper("Put"),
		"/v1/item/{id}",
		UpdateItemById,
	},
}
