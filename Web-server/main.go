package main

import (
	"fmt"
	"log"
	"net/http"
)

// func helloHandler(w http.ResponseWriter,r *http.Request){
// 	if r.URL.Path != "/hello"{			// checking the path
// 		http.Error(w,"404 page not found ",http.StatusNotFound) // status of the request
// 		return
// 	}
// 	if r.Method != "GET" { // methods cheks for the HTTP methods likes (GET,POST,PUT,DELETE)
// 		http.Error(w,"method is not supported",http.StatusNotFound)
// 		return
// 	}
// 	fmt.Fprintf(w,"Hello")  // Just writting Hello on the screen
// }

// func formHandler(w http.ResponseWriter,r *http.Request){
// 	if err := r.ParseForm();err != nil{    // parsing the form
// 		fmt.Fprintf(w,"ParseForm() err : %v",err)
// 		return
// 	}
// 	fmt.Fprintf(w,"POST REQUEST SUCCESSFULL \n")
// 	name :=r.FormValue("Name") 		//the string has to be equal to the label in the html file in the paranthesis
// 	address := r.FormValue("Address")   //to store the values that are entered in the form
// 	fmt.Fprintf(w,"Name = %s \n",name) //to print the stores values
// 	fmt.Fprintf(w,"Address = %s \n",address)
// }

// func main(){
// 	fileServer:= http.FileServer(http.Dir("./Static"))  //to create a server with a local file containing it to the folder of html and CSS
// 	http.Handle("/",fileServer) // to navigate through the file
// 	http.HandleFunc("/form",formHandler)  //regigtered the handler funtion for thr given path
// 	http.HandleFunc("/hello",helloHandler)

// 	fmt.Printf("Starting server at port 8080 \n")
// 	if err:= http.ListenAndServe(":8080",nil); err !=nil {		// creating a server at the given portno.		ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.
// 		log.Fatal(err)   //checking for the error
// 	}
// }

func main() {
	fileserver := http.FileServer(http.Dir("./Static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Starting server port at 8080 ")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "page not found 404", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Hello I actually made a web-server, Yayyyyyy")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v \n", err)
	}
	fmt.Fprintln(w, "Posted Successfull :")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name : %s \n", name)
	fmt.Fprintf(w, "Address : %s \n", address)
}
