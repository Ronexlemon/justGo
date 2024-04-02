package webrequests

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://atlas-ke.net"
func WebRequests(){
	welcome:="web requests"
	fmt.Println(welcome)
	res,err:=http.Get(url)
	checkNilError(err)
	fmt.Printf("type of response %T\n",res)
	
	

	defer res.Body.Close()  // your responsibility to close the response

	dataBytes,err:=io.ReadAll(res.Body)   // read all data from
	checkNilError(err)
	content := string(dataBytes) 
	fmt.Println(content) 
	  // convert bytes to a string


}

func checkNilError(err error){
	if err !=nil{
		panic(err)
	}

}