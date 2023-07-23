package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main (){
	

	godotenv.Load(".env")

	portString := os.Getenv("PORT")			// getting the port variable

	if portString == ""{ 								// checking if the port variable is empty if empty then -> 
		log.Fatal("Port variable not found")		//sending an error and immediately stopping the program
	}

	router := chi.NewRouter()				// creatng a new router

	srv := &http.Server{			//  we are connecting the router to the http server
		Handler: router,
		Addr: ":" + portString, 		// Addr optionally specifies the TCP address for the server to listen on
	}

	log.Printf("Server starting on port %v \n", portString)

	err := srv.ListenAndServe()			// ListenAndServe listens on the TCP network address srv.Addr and then calls Serve to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PORT : ", portString)
}