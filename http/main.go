package main

import (
	"fmt"
	"gohttp/pkg/server"
	"log"
	"net/http"
	"time"
)

//constants

/****
*const(
MethodGet ="GET"
MethodHead  ="Head"
MethodPost ="POST"
MethodPut  ="PUT"
MethodPatch ="PATCH"
MethodDelete ="DELETE"
MethodConnect ="CONNECT"
MethodOptions ="OPTIONS"
MethodTrace ="TRACE")
***************************************************

*******/

func main() {
	welcome := "Welcome to http package"
	simpleWebServer()
	fmt.Println(welcome)
}

func simpleWebServer() {
	mux := http.NewServeMux()
	address := ":8000"
	srv := server.New()

	mux.HandleFunc("/", srv.HandleIndex)
	mux.HandleFunc("/users/create", srv.HandleCreateUser)
	mux.HandleFunc("/user/{name}", srv.HandleGetUser)

	s := &http.Server{
		Addr:           address,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Start Server: %v", address)
	log.Fatal(s.ListenAndServe())
}