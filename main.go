package main

import (
	"log"
	"net/http"
	"wallets/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", controllers.RegisterHandler).Methods("POST")
	log.Println("Listen on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
