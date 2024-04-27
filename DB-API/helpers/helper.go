package helpers

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

var collection *mongo.Collection
const connectionString = "mongodb+srv://huge65702:huge65701@cluster0.9j3wobi.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0" //just for test

const dbName = "netflix"
const colName = "watchlist"
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

//insert one record

func InsertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.TODO(), movie)
	checkNil(err)
	fmt.Println("inserted one movie in db with id", inserted.InsertedID)

}

// update one record
func UpdateOneRecord(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	checkNil(err)
	fmt.Println("Successful update the watched field", result.ModifiedCount)
}

// delete one record
func DeleteOneRecord(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	checkNil(err)
	fmt.Println("delete successful", result)

}

// delete all
func DeleteAllRecords() {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})
	checkNil(err)
	fmt.Println("Number of movies deleted", deleteResult.DeletedCount)
}

//get all from mongo

func GetAllRecords() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}}) //return of a cursor
	checkNil(err)

	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		checkNil(err)
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies

}

func checkNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
