package main

import "net/http"

func handlerRediness(w http.ResponseWriter, r *http.Request) {		// adding a header for returning status okay if there is hit in the server
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}