package main

import (
	"fmt"
	"gomux/router"
	"log"
	"net/http"
)
func main(){
	welcome:="Go Mux "
	fmt.Println(welcome)

	log.Fatal(http.ListenAndServe(":4000",router.Router()))
}