package routesPost

import (
	"kdl_api_rest_internal/endpoint/post"

	"github.com/gorilla/mux"
)

func SetPostRoutes(router *mux.Router) {
	router.HandleFunc("/recordSimcard", post.RecordSimcard).Methods("POST")
}
