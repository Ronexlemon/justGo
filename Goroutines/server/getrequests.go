package server

import (
	"fmt"
	"io"
	"net/http"
)


func GetRequests(){
	welcome:="welcome to get requests"
	fmt.Println(welcome)
}

func PerformGetRequest(){
	const myUrl = "http://localhost:3000/get";
	response,err := http.Get(myUrl)
	checkNill(err)

	defer response.Body.Close()
	fmt.Println("Status code",response.StatusCode)
	fmt.Println("Content length", response.ContentLength)

	Content,_:=io.ReadAll(response.Body)
	fmt.Println(string(Content))

}

func checkNill(err error){
	if err != nil{
		panic(err)
	}

}