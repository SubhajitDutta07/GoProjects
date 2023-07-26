package main

import "net/http"

func handlerReadiness( w http.ResponseWriter, r *http.Request) {			// specific handler func to define a http handler in a way that go standard library expects
	respondWithJson(w,200,struct{}{})
	
}