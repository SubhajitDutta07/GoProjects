package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Quantity int `json:"quantity`
	Price float64 `json:"price"`
}

func get_Products(db *sql.DB)([]product, error){
	query := "select id, name , quantity , price From products"
	rows , err := db.Query(query)

	if err != nil {
		return nil , err
	}
	products := []product{}
	for rows.Next(){
		var p product
		err := rows.Scan(&p.ID,&p.Name,&p.Quantity,&p.Price)
		if err != nil {
			return nil ,err
		}
		products = append(products, p)
	}
	return products , nil 

}

func (p *product) getProduct( db *sql.DB) error {
	query := fmt.Sprintf("Select name, quantity, price from products where id=%v", p.ID)
	rows := db.QueryRow(query)
	err := rows.Scan(&p.Name,&p.Quantity,&p.Price)
	if err != nil {
		return err
	}
	return nil 
}

func (p *product) createProduct( db *sql.DB) error {
	query := fmt.Sprintf(" Insert into products(name,price,quantity) values('%v',%v,%v)", p.Name, p.Price, p.Quantity)
	result, err := db.Exec(query)
	if err != nil {
		return err
	}
	id ,err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = int(id)
	return nil 
}

func (p *product) updateProduct( db *sql.DB) error {
	query := fmt.Sprintf("update products set name='%v', price=%v, quantity=%v where id=%v",p.Name, p.Price,p.Quantity,p.ID)
	result, err := db.Exec(query)
	log.Println(result.RowsAffected())
	rowsAffecter,err := result.RowsAffected()
	if rowsAffecter == 0 {
		return errors.New("No such row exist")
	}

	return err
}

func (p *product) deleteProduct (db *sql.DB) error {
	query := fmt.Sprintf("delete from products where id=%v",p.ID)
	_,err := db.Exec(query)
	return err 
}