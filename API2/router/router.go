package router

import (
	"context"
	"fmt"

	"golangapis/config"
	"golangapis/model"
	"net/http"

	"github.com/gorilla/mux"
)


func ConnectDb() *config.EmployeeRepo {
	mongoClientTest := config.NewMongoClient()

	// defer statement should be placed after the potential error checks
	defer mongoClientTest.Disconnect(context.Background())

	// connect to collection
	collection := mongoClientTest.Database("company").Collection("employee")

	// return a pointer to EmployeeRepo
	return &config.EmployeeRepo{MongoCollection: collection}
}

func Insert(w http.ResponseWriter, r *http.Request){

	 params := mux.Vars(r)
	 employeeId:= params["employee_id"]
	 name:= params["name"]
	 department:= params["department"]

	 

	  emp:= &model.Employee{
		EmployeeID: employeeId,
		Name: name,
		Department: department,
	 }
	 fmt.Println("values are values",emp)


//   empRepo := ConnectDb()

//   result,err := empRepo.InsertEmployee()

 
}