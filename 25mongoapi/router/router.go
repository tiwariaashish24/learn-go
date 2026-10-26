package router

import (
	"mongodb/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	//router.Handlefunc().Methods("GET")

	router.HandleFunc("/api/movies", controller.GetMyAllMovie).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.MarkedAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovie).Methods("DELETE")

	return router
}
