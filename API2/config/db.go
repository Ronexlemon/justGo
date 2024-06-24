package config

import (
	"context"
	"fmt"
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

func (r *EmployeeRepo) FindAllEmployees()([]model.Employee,error){
	results,err := r.MongoCollection.Find(context.Background(),bson.D{})
	if err !=nil{
		return nil,err
	}

	var emps []model.Employee

	err = results.All(context.Background(),&emps)
	if err != nil {
		return nil, fmt.Errorf("results decode error %s",err.Error())
	}
	return emps,nil
	
}

func (r *EmployeeRepo) UpdateEmployeeById(empid string, updateEmp *model.Employee)(int64,error){
	result,err:= r.MongoCollection.UpdateByID(context.Background(),bson.D{{Key: "employee_id", Value: empid}},bson.D{{Key: "$set",Value: updateEmp}})
	if err != nil {
		return 0,err
		}
		return result.ModifiedCount,nil
}

func (r *EmployeeRepo) DeleteEmployeeById(empId string)(int64, error){

	result,err := r.MongoCollection.DeleteOne(context.Background(),bson.D{{Key: "employee_id", Value: empId}})
	if err != nil {
		return 0,err
		}
		return result.DeletedCount,nil
}

func (r *EmployeeRepo) DeleteAllEmployee()(int64, error){

	result,err := r.MongoCollection.DeleteOne(context.Background(),bson.D{})
	if err != nil {
		return 0,err
		}
		return result.DeletedCount,nil
}

func CheckNil(err error){
	if err != nil {
		panic(err)
		}
}