package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	Router chi.Router
	DB     *sql.DB
}

func (app *App) Initialize() error {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DBuser, DBpassword, DBname)
	var err error
	app.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	app.Router = chi.NewRouter()
	app.handleRoutes()
	return nil
}

func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func sendError(w http.ResponseWriter, statusCode int, err string) {
	err_msg := map[string]string{"error": err}
	sendResponse(w, statusCode, err_msg)
}

func (app *App) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := get_Products(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, products)

}

func (app *App) getProduct( w http.ResponseWriter, r *http.Request) {
	key , err := strconv.Atoi(chi.URLParam(r,"id"))
	if err != nil {
		sendError(w , http.StatusBadRequest, "Inavlid product id")
	}
	p := product{
		ID: key,
	}
	err = p.getProduct(app.DB)
	if err != nil {
		switch err {
			case sql.ErrNoRows: 
				sendError(w,http.StatusNotFound, "Product not found")
			default : 
				sendError(w,http.StatusInternalServerError, err.Error())
		}
		return
	}
	sendResponse(w,http.StatusOK, p)
}

func (app *App) createProduct( w http.ResponseWriter, r *http.Request) {
	var p product 
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		sendError(w , http.StatusBadRequest, "Inavlid request payload")
		return
	}
	err = p.createProduct(app.DB)
	if err != nil {
		sendError(w , http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w,http.StatusOK, p)
}

func (app *App) updateProduct( w http.ResponseWriter, r *http.Request) {
	key, err := strconv.Atoi(chi.URLParam(r,"id"))
	if err != nil {
		sendError(w,http.StatusBadRequest,"Invalid product ID")
		return
	}
	var p product

	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		sendError(w,http.StatusBadRequest,"Invalid request payload")
		return
	}
	p.ID = key
	err = p.updateProduct(app.DB)
	if err != nil {
		sendError(w,http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w,http.StatusOK,p)
}

func (app *App) deleteProduct(w http.ResponseWriter, r *http.Request){
	key , err := strconv.Atoi(chi.URLParam(r,"id"))
	if err != nil {
		sendError(w,http.StatusBadRequest,err.Error())
	}
	p:= product{ID: key}
	err = p.deleteProduct(app.DB)
	if err != nil {
		sendError(w,http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w,http.StatusOK,map[string]string{"result" : "successful deletion"})
}

func (app *App) handleRoutes() {
	app.Router.Get("/products", app.getProducts)
	app.Router.Get("/product/{id}", app.getProduct)
	app.Router.Post("/product", app.createProduct)
	app.Router.Put("/product/{id}", app.updateProduct)
	app.Router.Delete("/product/{id}",app.deleteProduct)
}
