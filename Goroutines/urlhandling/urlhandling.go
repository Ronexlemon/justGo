package urlhandling

import (
	"fmt"
	"net/url"
)

const myurl string = "http://lemon.dev:3000/learn?coursename=reactjs&paymentid=567878"
func UrlHandling(){
	welcome:="welcome to url handling"
	fmt.Println(welcome)
	fmt.Println(myurl)
	//parsing the url -> extracting information

	result,_:=url.Parse(myurl)
	fmt.Println(result.Scheme) // http or https
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	
	qParams := result.Query()
	fmt.Printf("the type os query params are %T\n",qParams)
	fmt.Println(qParams.Get( "coursename"))  // getting value of
	fmt.Println(qParams["paymentid"])   // getting all

	for _,val:= range qParams{
		fmt.Println("params is:", val)
	}

	partsOfUrl:= &url.URL{
		Scheme:"https",
		Host:  "www.google.com",
		Path: "/search",
		RawQuery: "q=Ronex",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)


}