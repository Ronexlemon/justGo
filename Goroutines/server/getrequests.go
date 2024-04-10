package server

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)


func GetRequests(){
	welcome:="welcome to get requests"
	fmt.Println(welcome)
}

func PerformGetRequest(){
	const myUrl = "https://atlas-ke.net/login";
	response,err := http.Get(myUrl)
	checkNill(err)

	defer response.Body.Close()
	fmt.Println("Status code",response.StatusCode)
	fmt.Println("Content length", response.ContentLength)

	var responseString strings.Builder

	Content,_:=io.ReadAll(response.Body)
	byteCount,_ := responseString.Write(Content)
	fmt.Println("byte count is",byteCount)
	fmt.Println(responseString.String())
	//fmt.Println(string(Content))

}

func checkNill(err error){
	if err != nil{
		panic(err)
	}

}