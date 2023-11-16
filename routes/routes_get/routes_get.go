package routesGet

import (
	"kdl_api_rest_internal/endpoint/get"

	"github.com/gorilla/mux"
)

func SetGetRoutes(router *mux.Router) {
	router.HandleFunc("/simcards", get.GetAllSimcards).Methods("GET")
}
