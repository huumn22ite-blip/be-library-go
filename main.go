package main

import (
	"be-library-go/db"
	
	"net/http"

	"github.com/gorilla/mux"
)
 func main(){
	db.Connect()
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	 w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":8080",router)
 }