package main

import (
	"log"
	"net/http"
	"sample-project/handlers"

	_ "sample-project/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")

	r.HandleFunc("/subjects", handlers.CreateSubject).Methods("POST")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Server started on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
