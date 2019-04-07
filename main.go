package main

import(
	"net/http"
	"os"
	"REST/api"
	"log"
	"github.com/gorilla/handlers"
)

func main(){
	//Get the PORT environment variable if it exists
	port := os.Getenv("PORT")


	if port == "" {
		//Set the port to localhost 8000 port
		port = "8000"
	}

	//Create a new router for redirecting requests to different handlers
	router := api.NewRouter()

	//Allow all types of origins for making request
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET","POST","PUT","DELETE"})
	//Start the server port number and router
	log.Fatal(http.ListenAndServe(":"+port,handlers.CORS(allowedOrigins,allowedMethods)(router)))

}


