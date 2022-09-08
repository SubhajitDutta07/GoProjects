package main

import (
	"fmt"
	"log"
	"github.com/gorilla/mux"

	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello Guys")
}

func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",hello).Methods("GET")
	myRouter.HandleFunc("/user",AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}",NewUser).Methods("POST")
	myRouter.HandleFunc("user/{name}",DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}",UpdateUser).Methods("PUT")
	fmt.Println("starting a port at :8080")
	log.Fatal(http.ListenAndServe(":8080",myRouter))
}

func main(){
	fmt.Println("Starting point")

	InitialMigration()

	handleRequests()
}