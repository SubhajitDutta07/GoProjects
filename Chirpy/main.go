package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type apiConfig struct {
	fileserverHits int
}

func main() {
	const filepathRoot = "."
	const port = "8080"

	apiCfg := apiConfig{
		fileserverHits: 0,
	}

	//mux := http.NewServeMux()
	//corMux := middlewareCors(mux)

	r := chi.NewRouter()
	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	r.Handle("/app", fsHandler)
	r.Handle("/app/*", fsHandler)
	r.Handle("/assets/logo.png", http.FileServer(http.Dir("assets/logo.png")))

	apiRouter := chi.NewRouter()
	apiRouter.Post("/validate_chirp", handlerChirpValidate)
	apiRouter.Get("/healthz", handlerRediness)
	r.Mount("/api", apiRouter) // added a new route for the api section

	adminRouter := chi.NewRouter()
	adminRouter.Get("/metrics", apiCfg.handlerMetrics)
	r.Mount("/admin", adminRouter) // added a new route for the admin section

	corsMux := middlewareCors(r)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}

func handlerChirpValidate(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	const maxChirp = 140
	if len(params.Body) > maxChirp {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}
	cleaned := getCleanedBody(params.Body, badWords)

	respondWithJson(w, http.StatusOK, returnVals{
		CleanedBody: cleaned,
	})

}

func getCleanedBody(body string,badWords map[string]struct{}) string {
	words := strings.Split(body, " ")
	for i,word := range words{
		lowerWord := strings.ToLower(word)
		if _,ok := badWords[lowerWord]; ok{
			words[i]="****"
		}
	}
	cleaned := strings.Join(words, " ")
	return cleaned
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error : %s", msg)
	}
	type errorResp struct {
		Error string `json:"error"`
	}

	respondWithJson(w, code, errorResp{
		Error: msg,
	})
}

func respondWithJson(w http.ResponseWriter, code int, payLoad interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payLoad)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}
