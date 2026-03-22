package routers

import (
	"be-library-go/handlers"

	"github.com/gorilla/mux"
)

func SetUpRouter() *mux.Router {
	router := mux.NewRouter()

	//cat

	router.HandleFunc("/categories", handlers.GetCategories).Methods("GET")
	router.HandleFunc("/categories", handlers.CreateCategories).Methods("POST")
	router.HandleFunc("/categories/{id}", handlers.UpdateCategories).Methods("PUT")
	router.HandleFunc("/categories/{id}", handlers.DeleteCategories).Methods("DELETE")

	//book

	router.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/books", handlers.CreateBooks).Methods("POST")
	router.HandleFunc("/books/{id}", handlers.UpdateBooks).Methods("PUT")
	router.HandleFunc("/books/{id}", handlers.DeleteBooks).Methods("DELETE")

	router.HandleFunc("/loans", handlers.GetLoans).Methods("GET")
	router.HandleFunc("/loans", handlers.CreateLoans).Methods("POST")
	router.HandleFunc("/loans/{id}", handlers.UpdateLoans).Methods("PUT")
	router.HandleFunc("/loans/{id}", handlers.DeleteLoans).Methods("DELETE")

	//members

	router.HandleFunc("/members", handlers.GetMembers).Methods("GET")
	router.HandleFunc("/members", handlers.CreateMembers).Methods("POST")
	router.HandleFunc("/members/{id}", handlers.UpdateMembers).Methods("PUT")
	router.HandleFunc("/members/{id}", handlers.DeleteMembers).Methods("DELETE")

	//staff

	router.HandleFunc("/staff", handlers.GetStaff).Methods("GET")
	router.HandleFunc("/staff", handlers.CreateStaff).Methods("POST")
	router.HandleFunc("/staff/{id}", handlers.UpdateStaff).Methods("PUT")
	router.HandleFunc("/staff/{id}", handlers.DeleteStaff).Methods("DELETE")
	return router
}
