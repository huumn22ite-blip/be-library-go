                                                                        
package main

import (
	"fmt"
	"be-library-go/db"
	"be-library-go/routers"
	"net/http"


)
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
func main() {
	db.Connect()

	router := routers.SetUpRouter()


	fmt.Println("Server running on port 8080")

http.ListenAndServe(":8080", enableCORS(router))
}

