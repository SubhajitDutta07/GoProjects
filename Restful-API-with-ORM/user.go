package main

import (
	"fmt"
	"net/http"

	"github.com/go-yaml/yaml"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB		// defining our database connection
 
var err error

type User struct{
	gorm.Model		//to let the gorm know that,this is the gorm model of the database
	Name string 
	Email string

}

func InitialMigration(){
	
	if db, err = gorm.Open("sqlite", "test.db"); err != nil{		//to open sqlite type database
		fmt.Println(err.Error())
		panic("Failed to connect to the database ")
	}

	defer db.Close()

	db.AutoMigrate(&User{}) //to automatically migrate 

}


func AllUsers( w http.ResponseWriter , r *http.Request){ 
	//test.db is a database file 
	if db,err = gorm.Open("splite","test.db"); err != nil{  // this function will return all of the functions in our database 
		panic("Could not connect to the database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)  // it finds records that matchies given condition
	yaml.NewEncoder(w).Encode(users)

}

func NewUser(w http.ResponseWriter, r *http.Request){
	if db,err = gorm.Open("sqlite","test.db"); err != nil{
		panic("Could not open the database file ")
	}
	defer db.Close()
	// now to need to capture the path parameters of the request made to our API
	vars := mux.Vars(r)
	name := vars["name"] 	// parses the path parameter for the name
	email:= vars["email"]

	db.Create(&User{Name: name , Email: email})

	fmt.Fprintf(w,"New user has been created")

}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	if db,err = gorm.Open("sqlite","test.db"); err != nil{
		panic("Could not open the database file ")
	}
	defer db.Close()

	vars :=mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?",name).Find(&user)  // where is the user located and to find them 
	db.Delete(&user) // delete the user

	fmt.Fprintf(w,"User successfully deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	if db,err = gorm.Open("sqlite","test.db"); err != nil{
		panic("Could not open the database file ")
	}
	defer db.Close()
	vars:= mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name=?",name).Find(&user)

	user.Email= email
	db.Save(&user)
	fmt.Fprintf(w,"Successfully Updated user")


}	