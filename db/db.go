package db

import (
	"database/sql"
	"log"
	

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
func Connect(){
	var err error
DB, err = sql.Open("mysql", "root:dgstFAeTojWneuiaOEFsjfKUvVqHTuey@tcp(hopper.proxy.rlwy.net:50741)/library")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("keet noi ok ")
}
