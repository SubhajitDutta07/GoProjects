package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var movies[]Movie

func getMovies(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type","application/json")	// setting the content type to yaml so that golang can convert the incoming data in yaml format
	json.NewEncoder(w).Encode(movies) // encoding the slice into yaml

}

func deleteMovie(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Conetent-Type","application/json")
	params := mux.Vars(r)		// to get something from the user
	for k,v := range movies {

		if v.ID == params["id"] {
			movies = append(movies[:k],movies[k+1:]...) 	//to delete a movie
			break
		}
	}
	json.NewEncoder(w).Encode(movies)		// again encoding the slice into yaml format
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:= mux.Vars(r)
	for _,v:=range movies{
		if v.ID == params["id"]{
			json.NewEncoder(w).Encode(v)
			return
		}
	}

}

func createMovies( w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_=json.NewDecoder(r.Body).Decode(&movie) 		// addind a new movie with the synatx of json
	movie.ID = strconv.Itoa(rand.Intn(10000)) 	//getting a random no.
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for k,v := range movies {
		if v.ID == params["id"]{
			movies = append(movies[:k],movies[k+1:]...  )
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]   	// same id as the deleted movie
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}


func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1",Isbn:"4768", Title: "Movie One", Director: &Director{FirstName: "Lenarnd", LastName: "Cooper"} })
	movies = append(movies, Movie{ID: "2",Isbn: "8565", Title: "Movie Two", Director: &Director{FirstName: "Sheldon", LastName: "Hoafsteder"}})
	movies = append(movies, Movie{ID: "3", Isbn: "7654", Title: "Movie Three", Director: &Director{FirstName: "Bernadette", LastName: "Rostankatski"}})
	r.HandleFunc("/movies",getMovies).Methods("GET")		//creted a movies server with the method "GET"
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movijsonid}",deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080",r))		// creating a live server and checking for any error, if there is any error we will exit the applicatiton
}