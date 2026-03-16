package db

import (
	"database/sql"
	"log"
	 _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
func Connect(){
	var err error
	DB,err = sql.Open(
		"mysql", 	"root:123456@tcp(127.0.0.1:3306)/library",
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("keet noi ok ")
}