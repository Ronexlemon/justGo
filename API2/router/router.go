package router

import (
	// "context"
	"encoding/json"
	"fmt"

	"golangapis/config"
	"golangapis/model"
	"net/http"
	// "github.com/gorilla/mux"
)

func ConnectDb() *config.EmployeeRepo {
	mongoClientTest := config.NewMongoClient()

	// defer statement should be placed after the potential error checks
	// defer mongoClientTest.Disconnect(context.Background())

	// connect to collection
	collection := mongoClientTest.Database("company").Collection("employee")

	// return a pointer to EmployeeRepo
	return &config.EmployeeRepo{MongoCollection: collection}
}

func Insert(w http.ResponseWriter, r *http.Request) {
	var emp model.Employee
	

	err := json.NewDecoder(r.Body).Decode(&emp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	   empRepo := ConnectDb()

	   result,err := empRepo.InsertEmployee(&emp)
	   if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return}
		
		fmt.Println(result)
		
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User Created successful")


		
}

func Employees(w http.ResponseWriter, r *http.Request){
	empRepo := ConnectDb()

	employees, err := empRepo.FindAllEmployees()
	if err !=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(employees)
}