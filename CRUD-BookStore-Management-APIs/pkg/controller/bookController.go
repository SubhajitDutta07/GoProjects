package controller

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"CRUD-BookStore-Management-APIs/pkg/utils"
	"CRUD-BookStore-Management-APIs/pkg/models"
)

var NewBook modules.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks:= modules.GetAllBooks()
	//marshal returns slice of byte and an error
	res, _ :=json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	// write helps us the send something to the frontend
	// here json of the NewBooks
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookid"]
	ID, err:= strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	bookDetails, _:=modules.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &modules.Book{}
	utils.ParseBody(r,CreateBook)
	b:= CreateBook.CreateBook()
	res,_ := json.Marshal(b)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars:= mux.Vars(r)
	bookID :=vars["bookid"]
	ID, err := strconv.ParseInt(bookID,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	book := modules.DeleteBook(ID)
	res, _ :=json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &modules.Book{}
	utils.ParseBody(r,updateBook)
	vars := mux.Vars(r)
	bookID := vars["bookid"]
	ID,err := strconv.ParseInt(bookID,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	bookDetails, db := modules.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication= updateBook.Publication
	}
	// save in our database
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}