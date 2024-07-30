package main

import (
	"fmt"
	"io"
	"net/http"
	
)

func main(){
	GetFunc()

	

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