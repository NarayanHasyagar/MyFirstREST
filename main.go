package main;

import (

	"net/http"
	
	"os"
	"api"
)

func main(){
	//Get the PORT environment variable if it exists
	port := os.GetEnv("PORT")


	if port == "" {
		//Set the port to localhost 8000 port
		port = "8000"
	}

	//Create a new router for redirecting requests to different handlers
	router = handlers.NewRouter()

	//Start the server port number and router
	log.Fatal(http.ListenAndServe(":"+port,router))

}


