package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"goapi/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://huge65702:huge65701@cluster0.9j3wobi.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0" //just for test

const dbName = "netflix"
const colName = "watchlist"

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

//Controllers

func GetAllMovies(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	allMovies:=getAllRecords()
	json.NewEncoder(w).Encode(allMovies)
}

//delete one record

func DeleteAMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	params:= mux.Vars(r)

	deleteOneRecord(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

//function deletAll
func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	deleteAllRecords()
	json.NewEncoder(w).Encode("All movies deleted")
}

//update on record

func MarkMovieAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")
	
	params:=mux.Vars(r)
	updateOneRecord(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

//function insert
func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")
	
	var movie model.Netflix

	_=json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func checkNil(err error){
	if err !=nil{
		log.Fatal(err)
	}
}

