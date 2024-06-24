package main

import (
	"fmt"
	"golangapis/config"
)


func main(){
	welcome := "welcome to api development"

	fmt.Println(welcome)
	config.NewMongoClient()
}