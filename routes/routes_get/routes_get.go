package routesGet

import (
	"kdl_api_rest_internal/endpoint/get"

	"github.com/gorilla/mux"
)

func SetGetRoutes(router *mux.Router) {
	router.HandleFunc("/simcards", get.GetAllSimcards).Methods("GET")
	router.HandleFunc("/listIccids", get.GetListIccids).Methods("GET")
	router.HandleFunc("/stock", get.GetAllStock).Methods("GET")
	router.HandleFunc("/logs", get.GetAllLogs).Methods("GET")
}
