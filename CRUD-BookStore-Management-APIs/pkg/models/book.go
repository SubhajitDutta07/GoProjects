package modules

import (
	"github.com/jinzhu/gorm"
	"CRUD-BookStore-Management-APIs/pkg/config"
)

var db *gorm.DB
// book data structure
type Book struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	Author	string `json:"author"`
	Publication string `json:"publication"`
}
// initializing the database
func init (){
	//connectting with the database
	config.Connect()
	//returning the database variable as in the pointer to the database
	db = config.GetDB()
	// to migrate the data into an empty structure
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	//NewRecord check if value's primary key is blank
	db.NewRecord(b)
	//Create insert the value into database
	db.Create(&b)
	return b
}

func GetAllBooks () []Book{
	var Books []Book
	//Find find records that match given conditions
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64)( *Book, *gorm.DB){
	var getBook Book
	//Where return a new relation, filter records with given conditions, accepts `map`, `struct` or `string` as conditions
	db = db.Where("ID=?",Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?",ID).Delete(book)
	return book
}