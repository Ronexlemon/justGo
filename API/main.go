package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct{
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int   `json:"price"`
	Author *Author `json:"author"`

}

type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`

}


//fake db

var courses []Course
//helper/middleware functions
func (c *Course) IsEmpty()bool{

	return c.CourseId == "" && c.CourseName == ""


}

/*
*action plan ,create,delete,search
**/
func main() {
	welcome := "hellow there welcome to api junkie"
	fmt.Println(welcome)
}


//controllers -file

//serve Home route

func serveHome(w http.ResponseWriter,r *http.Request){
			w.Write([]byte("<h1>welcome to API by lemonr</h1>"))

}

//get all

func getAllCourses(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type","application/json")

	json.NewEncoder(w).Encode(courses)
}


func getOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("get one course")
	w.Header().Set("Content-Type","application/json")
	//grab id from request

	params := mux.Vars(r)
	//loop through courses,find matching id and return the response

	for _,course := range courses{

		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found ")
	return
}

