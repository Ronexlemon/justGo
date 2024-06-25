package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golangapis/router"

	"github.com/gorilla/mux"
)


func main(){
	welcome := "welcome to api development"
	route := mux.NewRouter()

	route.HandleFunc("/create",router.Insert).Methods("POST")
	route.HandleFunc("/employees",router.Employees).Methods("GET")
	srv := &http.Server{
		Handler:      route,
		Addr:         "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println(welcome)
	log.Fatal(srv.ListenAndServe())
	
}