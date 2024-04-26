package controller

import (
	"context"
	"fmt"
	"goapi/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = " mongodb+srv://huge65702:huge65701@cluster0.9j3wobi.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0" //just for test

const dbName = "netflix"
const colName = "watvhlist"

var collection *mongo.Collection

//connect with mongodb

func init() {
//client options
clientOption:= options.Client().ApplyURI(connectionString)

//connect to db

client,err:=mongo.Connect(context.TODO(),clientOption)
checkNil(err)
fmt.Println("mongodb connection succesful")

//create collection
collection = client.Database(dbName).Collection(colName)
//for collection check
fmt.Println("collection instance is ready yollow!!")


}


//mongodb helpers - file

//insert one record

func insertOneMovie(movie model.Netflix){
	inserted,err:=collection.InsertOne(context.TODO(),movie)
	checkNil(err)
	fmt.Println("inserted one movie in db with id",inserted.InsertedID)

}
//update one record
func updateOneRecord(movieId string){
	id,_:=primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	update:= bson.M{"$set":bson.M{"watched":true}}

	result,err:=collection.UpdateOne(context.Background(),filter,update)
	checkNil(err)
	fmt.Println("Successful update the watched field",result.ModifiedCount)
}


func checkNil(err error){
	if err !=nil{
		log.Fatal(err)
	}
}

