package main


import (
	"net/http"
	"fmt"
)


func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) { // adding a header for the no. of hits in the server
	w.Header().Add("Content-Type", "text/plain ; charset=utf-8 ")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits : %d \n", cfg.fileserverHits)))
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler { // func for adding the no. of time the server is visited

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits++
		next.ServeHTTP(w, r)
	})
}