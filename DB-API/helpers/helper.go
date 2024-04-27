package helpers

import (
	"context"
	"fmt"
	"goapi/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


var collection *mongo.Collection
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

//delete one record
func deleteOneRecord(movieId string){
	id,_:= primitive.ObjectIDFromHex(movieId)

	filter:= bson.M{"_id":id}
	result,err:= collection.DeleteOne(context.Background(),filter)
	checkNil(err)
	fmt.Println("delete successful",result)

	
}
//delete all
func deleteAllRecords(){
	
	deleteResult,err :=collection.DeleteMany(context.Background(),bson.D{{}})
	checkNil(err)
	fmt.Println("Number of movies deleted",deleteResult.DeletedCount)
}

//get all from mongo

func getAllRecords()[]primitive.M{
	cur,err:=collection.Find(context.Background(),bson.D{{}}) //return of a cursor
	checkNil(err)
	
	var movies []primitive.M
	for cur.Next(context.Background()){
		var movie bson.M
		err:=cur.Decode(&movie)
		checkNil(err)
		movies=append(movies,movie)
	}

	defer cur.Close(context.Background())
	return movies
	


}


func checkNil(err error){
	if err !=nil{
		log.Fatal(err)
	}
}