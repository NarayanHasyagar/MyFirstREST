package api

import(
	"net/http"
	"github.com/gorilla/mux"
)

//Create an empty structure pointer to refer to methods
var controller = &Controller{Repository: Repository{}}

//Structure for storing routing parameters
type Route struct
{

	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

//Structure slice to store information
type Routes []Route

var routes = Routes{
	Route{
		"Welcome",
		"GET",
		"/",
		controller.Welcome,
	},
	Route{
		"ViewAllProduct",
		"GET",
		"/Product",
		controller.ViewAllProduct,

	},
	Route{
		"AddProduct",
		"POST",
		"/Product/Add",
		controller.AddProduct,
	},
	Route{
		"ViewProduct",
		"GET",
		"/Product/{id}",
		controller.ViewProduct,
	}}
//Implement Router method to handle the requests
func NewRouter() *mux.Router{

	//Create a mux router
	router := mux.NewRouter()
	//Assign handler functions for different paths
	for _,route := range routes{
		router.Name(route.Name).Methods(route.Method).Path(route.Pattern).HandlerFunc(route.HandlerFunc)
	}
	return router
}