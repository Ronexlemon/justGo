package main

import (
	"fmt"
	"net/http"
)

func main(){

	res, err := http.Get("https://github.com/RonexLemon")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)

}

