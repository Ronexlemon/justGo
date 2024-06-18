package config

import (
	"context"
	"golangapis/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var value string = "mongodb+srv://huge65702:huge65702@cluster0.3s02tjq.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

type EmployeeRepo struct{
	MongoCollection *mongo.Collection
}

func  (r *EmployeeRepo) InsertEmployee(emp *model.Employee)(interface{},error){
	result, err := r.MongoCollection.InsertOne(context.Background(),emp)
	if err != nil {
		return nil,err
		}

		return result,nil
}

func (r *EmployeeRepo) FindEmployeeById(empId string)(*model.Employee,error){
	var emp  model.Employee
	err:= r.MongoCollection.FindOne(context.Background(),bson.D{{Key: "employee_id", Value:empId}}).Decode(&emp)
	if err != nil {
		return nil,err
		}
		return &emp,nil
}


func CheckNil(err error){
	if err != nil {
		panic(err)
		}
}