package routesGet

import (
	"kdl_api_rest_internal/endpoint/get"

	"github.com/gorilla/mux"
)

func SetGetRoutes(router *mux.Router) {
	router.HandleFunc("/simcards", get.GetAllSimcards).Methods("GET")
<<<<<<< HEAD
	router.HandleFunc("/simucs", get.GetAllSimucs).Methods("GET")
=======
	router.HandleFunc("/listIccids", get.GetListIccids).Methods("GET")
	router.HandleFunc("/stock", get.GetAllStock).Methods("GET")
	router.HandleFunc("/logs", get.GetAllLogs).Methods("GET")
>>>>>>> c122bb3b4c87e543ccda90ca11c6360bc5c80e01
}
