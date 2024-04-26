package router

import (
	"goapi/controller"

	"github.com/gorilla/mux"
)





func Router()*mux.Router{
	router:=mux.NewRouter()

	router.HandleFunc("/api/movies",controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/deleteAll",controller.DeleteAllMovies).Methods("DELETE")
	router.HandleFunc("/api/movie",controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}",controller.MarkMovieAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie",controller.DeleteAMovie).Methods("DELETE")

	return router
}

