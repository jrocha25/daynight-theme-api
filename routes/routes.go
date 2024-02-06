package routes

import (
	"daynight-theme.dev/api/handlers"
	"github.com/gorilla/mux"
)

func MainRouter() *mux.Router {
	r := mux.NewRouter()
	subRouter := r.PathPrefix("/v1").Subrouter()

	subRouter.HandleFunc("/times", handlers.GetTimesHandler).Methods("GET")
	subRouter.HandleFunc("/countries", handlers.GetAllCountriesHandler).Methods("GET")

	return r
}
