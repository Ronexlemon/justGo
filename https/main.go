package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// your code here
	router:= mux.NewRouter()
	
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx := context.WithValue(r.Context(), "UserId", "12345")

		id := ctx.Value("UserId")
		//return http response to user
		fmt.Println(id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(id)

	}) 
	router.HandleFunc("/Data",GetMyData).Methods("GET")
	router.HandleFunc("/{id}",passIt)
	fmt.Println("Listening server .....")
	
	
	log.Fatal(http.ListenAndServe(":8080", router))
}


func GetMyData(w http.ResponseWriter, r *http.Request){
	
	res,err:= http.Get("https://github.com/RonexLemon")
	if err !=nil{
		log.Println(err)
		return
	}
	defer res.Body.Close()
	body,err:= io.ReadAll(res.Body)
	if err!=nil{
		log.Println(err)
		return
		}
		json.NewEncoder(w).Encode(map[string]string{
			"Data": string(body),
		})
}

//passing parameter
func passIt(w http.ResponseWriter, r *http.Request){
	val := mux.Vars(r)
	id:= val["id"]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"ID": id,
	})

}