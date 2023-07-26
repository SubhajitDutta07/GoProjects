package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/cors"
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

	router.Use(cors.Handler(cors.Options {						// this is so that people can make requestes from a browser
		AllowedOrigins: []string{"https://*", "http://*"},					// this configuration is for our handler to send bunch of extra http or https headers in our responses 
		AllowedMethods : []string {"GET", "POST", "PUT","DELETE","OPTIONS"},
		AllowedHeaders : []string{"*"},
		ExposedHeaders : []string{"Link"},
		AllowCredentials : false,
		MaxAge : 300,
	}))

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