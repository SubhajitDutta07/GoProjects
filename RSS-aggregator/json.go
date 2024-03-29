package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface {}){
	dta , err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal JSON response : %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/JSON")
	w.WriteHeader(code)
	w.Write(dta)

}