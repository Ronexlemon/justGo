package main

import (
	"fmt"
	"goapi/router"
	"log"
	"net/http"
)


func main(){
	welcome:= "Welcome to mongo api go"
	r:= router.Router()
	fmt.Println(welcome)
	 log.Fatal(http.ListenAndServe(":4000",r)) 
	 fmt.Println("listening at port 4000 ....")

	
}