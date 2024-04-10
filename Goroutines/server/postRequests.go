package server

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

var url = "http://localhost:3000/post"

func PerformPostRequests(){
	welcome:= "welcome to post requests"
	fmt.Println(welcome)

	//fake json payload

	requestsBody := strings.NewReader(`
	{"county":"Nairobi",
	"female":5000,
	"male":5000,
	"population":10000}
	`)

	res,err:=http.Post(url,"application/json",requestsBody)
	checkNilError(err)

	defer res.Body.Close()
	var responseString strings.Builder

	content,_:= io.ReadAll(res.Body)
	_,err = responseString.Write(content)
	checkNilError(err)
	
	fmt.Println(responseString.String())


}

func checkNilError(err error){
	if err != nil{
		panic(err)
	}
}