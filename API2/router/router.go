package router

import (
	"context"
	
	"golangapis/config"
	"net/http"
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

  //empRepo := ConnectDb()

 
}