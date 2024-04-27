package controller

import (
	
	"encoding/json"
	
	"goapi/model"
	"goapi/helpers"
	
	
	"net/http"

	"github.com/gorilla/mux"
	
)



//Controllers

func GetAllMovies(w http.ResponseWriter, r *http.Request){
	

	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	allMovies:=helpers.GetAllRecords()
	json.NewEncoder(w).Encode(allMovies)
}

//delete one record

func DeleteAMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	params:= mux.Vars(r)

	helpers.DeleteOneRecord(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

//function deletAll
func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	helpers.DeleteAllRecords()
	json.NewEncoder(w).Encode("All movies deleted")
}

//update on record

func MarkMovieAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")
	
	params:=mux.Vars(r)
	helpers.UpdateOneRecord(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

//function insert
func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")
	
	var movie model.Netflix

	_=json.NewDecoder(r.Body).Decode(&movie)
	helpers.InsertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}


