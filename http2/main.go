package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main(){
	// GetFunc()
	// fmt.Println("Calling the second function")
	// Client()
	Transport()

	

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
		Transport: http.DefaultTransport,
		
	}
	res,_ :=client.Get("https://celoafricadao.xyz")

	data,_ :=io.ReadAll(res.Body) 
	fmt.Println(data)

}

func Transport(){
	tr:= &http.Transport{
		DisableCompression: true,
		MaxIdleConns: 10,
		IdleConnTimeout: 30 *time.Second,

	}
	client:= &http.Client{Transport: tr}

	res,_:= client.Get("https://tracker.celoafricadao.xyz")

	result,err:= io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))
}