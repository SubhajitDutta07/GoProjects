package main

import (
	"log"
	"net/http"

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
	apiRouter.Get("/healthz", handlerRediness)
	r.Mount("/api", apiRouter)					// added a new route for the api section

	adminRouter := chi.NewRouter()
	adminRouter.Get("/metrics", apiCfg.handlerMetrics)
	r.Mount("/admin",adminRouter)				// added a new route for the admin section 
	

	corsMux := middlewareCors(r)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
