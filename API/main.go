package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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

	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseId == ""


}

/*
*action plan ,create,delete,search
**/
func main() {
	welcome := "hellow there welcome to api junkie"
	fmt.Println(welcome)
	//init router

	router:= mux.NewRouter()

	//seeding data;
	courses = append(courses, Course{CourseId: "2",CourseName: "Golang",CoursePrice: 300,Author: &Author{Fullname: "John doe",Website: "lemonr.io"}})
	courses = append(courses, Course{CourseId: "3",CourseName: "Nextjs",CoursePrice: 500,Author: &Author{Fullname: "hollow kale",Website: "go.io"}})

	//routing
	router.HandleFunc("/",serveHome).Methods("GET")
	router.HandleFunc("/courses",getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	router.HandleFunc("/course",createOneCourse).Methods("POST")
	router.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	router.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")
	
	

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000",router))
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


func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create one course")
	w.Header().Set("Content-Type","application/json")

	//what if body is empty

	if(r.Body) == nil{
		json.NewEncoder(w).Encode("Please Send some data")
		
	}

	//what about the data that is sent like {}

	var course Course

	_=json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
	//generate unique id


	rand.Seed(time.Now().UnixNano())

	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)

	return

	

	
}


//update

func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update one course")
	w.Header().Set("Content-Type","application/json")

	//first grab id from request

	params := mux.Vars(r)
	//loop, id, remove, add with my ID
	for index,course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_=json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}


	}
	json.NewEncoder(w).Encode("No course found with given id")
	

}

//delete

func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update one course")
	w.Header().Set("Content-Type","application/json")

	//first grab id from request

	params := mux.Vars(r)
	//loop, id, remove, add with my ID
	for index,course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			
			break
			
			
		}


	}
	json.NewEncoder(w).Encode("delete successful")
	

}