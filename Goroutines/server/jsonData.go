package server

import (
	"encoding/json"
	"fmt"
	"strings"
)

type course struct {
	Name     string   `json:"course_name"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

type courseA struct {
	Name     string   `json:"course_name"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}
type courseB struct {
	Name     string   `json:"course_name"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

type courseBB struct {
	Name     string   `json:"course_name"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

type courseBBC struct {
	Name     string   `json:"course_name"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

type courseBBCe struct {
	Name     string   `json:"course_name"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func JsonData() {
	welcome := "welcome to json data"
	fmt.Println(welcome)
}

//encode json

func EncodeJson() {
	lemonrCourses := []course{{Name: "go lang", Price: 200, Platform: "lemonr.io", Password: "Lemonr9090$", Tags: []string{"webdev", "js"}},
		{Name: "Java", Price: 250, Platform: "le.io", Password: "Le9090$", Tags: []string{"ful-satck", "ts"}},
		{Name: "Solidity", Price: 270, Platform: "lem.io", Password: "Lemonr9600$", Tags: []string{"backend", "ts"}},
		{Name: "Python", Price: 280, Platform: "ronex.com", Password: "ronex67$", Tags: nil}}

	//package data as json data
	var responseString strings.Builder
	//finalJson, err:= json.Marshal(lemonrCourses)
	finalJson, err := json.MarshalIndent(lemonrCourses, "", "\t")
	checkNilError(err)
	responseString.Write(finalJson)

	fmt.Println(responseString.String())

}

func DecodeJson(){
	jsonData:= []byte(`
	{
		"course_name": "go lang",
		"price": 200,
		"website": "lemonr.io",
		"tags": ["webdev","js"]
}	
	`)

	var lemonrCourses course

	checkValid :=json.Valid(jsonData)
	if checkValid{
		fmt.Println("Json data wa valid")
		json.Unmarshal(jsonData, &lemonrCourses)
		fmt.Printf("%#v\n",lemonrCourses)
	}else{
		fmt.Println("All json was not valid")
	}

	//some cases where you just want to add data to key value

	var myCourses map[string]interface{}
	json.Unmarshal(jsonData, &myCourses)
	fmt.Printf("%#v\n",myCourses)

}