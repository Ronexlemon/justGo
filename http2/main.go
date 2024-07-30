package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main(){
	GetFunc()
	fmt.Println("Calling the second function")
	Client()

	

}


func GetFunc(){
	res, err := http.Get("https://github.com/RonexLemon")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	fmt.Println(body)

	
}

func Client(){
	client:= http.Client{
		Timeout: 10 * time.Second,
		
	}
	res,_ :=client.Get("https://celoafricadao.xyz")

	data,_ :=io.ReadAll(res.Body) 
	fmt.Println(data)

}