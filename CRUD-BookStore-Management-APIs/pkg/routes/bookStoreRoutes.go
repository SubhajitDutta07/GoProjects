package routes

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"CRUD-BookStore-Management-APIs/pkg/controller"
)

var RegisterBookStoreRoutes = func ()  {
	r := mux.NewRouter().StrictSlash(true)
	fmt.Println("statring port at :8080")
	log.Fatal(http.ListenAndServe(":8080",r))

	http.Handle("/",r)
	r.HandleFunc("/book/",controller.CreateBook).Methods("POST")
	r.HandleFunc("/book/",controller.GetBook).Methods("GET")
	r.HandleFunc("/book/{bookid}",controller.GetBookById).Methods("GET")
	r.HandleFunc("/book/{bookid}",controller.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{bookid}",controller.DeleteBook).Methods("DELETE")

}