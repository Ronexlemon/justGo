package main

import (
	"fmt"
	"net/http"
	// "time"
)

func main(){
	welcome:="Welcome to go concurrency"
	fmt.Println(welcome)
	getStatusCode("https://github.com/RonexLemon")
	
}

// func greeter(s string){
// 	for i:=0; i< 6;i++{
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string){
	res,err:=http.Get(endpoint)
	checkNill(err)

	fmt.Println(res.Status)
}

func checkNill(err error){
	if err!=nil{
		panic(err)
	}
}