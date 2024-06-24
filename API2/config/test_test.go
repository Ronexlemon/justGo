package config

import (
	"context"
	
	"golangapis/model"
	"testing"

	"github.com/google/uuid"
)


func TestMongoOperations(t *testing.T){
	// Initialize the MongoDB connection
	
	mongoClientTest := NewMongoClient()

	defer mongoClientTest.Disconnect(context.Background())

	//dummy data
	empl1:= uuid.New().String()
	//empl2:= uuid.New().String()


	// coonect to collection
	collection := mongoClientTest.Database("company").Collection("employee")

	empRepo := EmployeeRepo{MongoCollection: collection}

	//insert employee 1 data

	t.Run("insert employee 1 ",func(t *testing.T){
		emp1:= model.Employee{
			EmployeeID: empl1,
			Name: "John",
			Department: "Computer science",
		}
		result,err :=empRepo.InsertEmployee(&emp1)
		if err != nil {
			t.Errorf("Error inserting employee 1: %v", err)
			}
			t.Log("insert user succssful",result)

	})


}