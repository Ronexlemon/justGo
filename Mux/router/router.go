package router

import (
	"encoding/json"
	"gomux/models"
	"net/http"

	"github.com/gorilla/mux"
)




func Router()*mux.Router{
	router:= mux.NewRouter()
	models.Stores = append(models.Stores, models.Store{StoreId: "1",StoreName: "Kilimall",StoreAddress: "Kenyataa avenue",StoreProducts: &models.Product{ProductId: "100",ProductName: "Nails",ProductPrice: 100}})
	models.Stores = append(models.Stores, models.Store{StoreId: "2",StoreName: "Jumia",StoreAddress: "Tomboya",StoreProducts: &models.Product{ProductId: "200",ProductName: "Polish",ProductPrice: 106}})
	models.Stores = append(models.Stores, models.Store{StoreId: "3",StoreName: "Choko",StoreAddress: "Kimathi street",StoreProducts: &models.Product{ProductId: "300",ProductName: "Gem",ProductPrice: 1050}})

	
	router.HandleFunc("/stores",getAllStores).Methods("GET")
	router.HandleFunc("/store/{id}",getOnestore).Methods("GET")

	return router

	
}


func getAllStores(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(models.Stores)

}

//get on store

func getOnestore(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	params:=mux.Vars(r)

	for _,item:=range models.Stores{
		if item.StoreId == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("Store not found")
	


}