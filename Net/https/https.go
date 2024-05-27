package https

import (
	"fmt"
	"net/http"
)

//Package http provides HTTP client and server implementation
var URL string = "https://example.com/"


func GETHeadPostPostForm(){

	resp,err := http.Get(URL)
	
	checkNil(err)
	fmt.Println(resp.Status)

	defer resp.Body.Close()

}

//create a client



func checkNil(err error){
	if err !=nil{
		panic(err)
	}
}