package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-yaml/yaml"

	"github.com/gorilla/mux"
	// it actually help us to identify what verbs do we need to access the endpoints and a far simpler fashion
	// we can add the methods as well
)									

type Article struct{
	Title string `yaml:"Title"`
	Describtion string `yaml:"Describtion"`
	Content string `yaml:"Content"`

}

var Articles []Article

func allArticles (w http.ResponseWriter, r *http.Request){
	Articles = append(Articles ,Article{ Title: "Title One", Describtion: "Describtion one", Content: "Content one"})
	fmt.Println("Endpoint Hit : All articles Endpoint")
	yaml.NewEncoder(w).Encode(Articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Test Post endpoint worked ")
}

func homePage( w http.ResponseWriter, r *http.Request ){
	fmt.Fprintf(w,"HomePage end point")
}

func handleRequest(){

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",homePage)
	myRouter.HandleFunc("/articles",allArticles).Methods("GET")
	myRouter.HandleFunc("/articles",testPostArticles).Methods("POST")
	fmt.Println("Starting server at port 8080 : ")
	log.Fatal(http.ListenAndServe(":8080",myRouter))
}
func main(){
	handleRequest()
}