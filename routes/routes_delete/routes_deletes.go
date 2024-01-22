package routesDelete

import (
	delete "kdl_api_rest_internal/endpoint/delete"

	"github.com/gorilla/mux"
)

func SetDeleteRoutes(router *mux.Router) {
	router.HandleFunc("/delete", delete.DeleteSimcards).Methods("DELETE")
	router.HandleFunc("/deleteStock", delete.DeleteSimcardsStock).Methods("DELETE")
}
