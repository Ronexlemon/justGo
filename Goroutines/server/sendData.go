package server

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func PerformPostFormRequets() {
	const myurl = "http://localhost:8000/postForm"

	data := url.Values{}
	data.Add("first_name", "John")
	data.Add("last_name", "John")
	data.Add("email","johndoe@gmail.com")

	response,err :=http.PostForm(myUrl,data)
	checkNilError(err)

	defer response.Body.Close()
	var responseString strings.Builder

	content,err := io.ReadAll(response.Body)
	checkNilError(err)
	_,err= responseString.Write(content)
	checkNilError(err)
	fmt.Println("content is", responseString.String())


}
