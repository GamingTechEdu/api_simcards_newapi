package routesPost

import (
	"kdl_api_rest_internal/endpoint/post"

	"github.com/gorilla/mux"
)

func SetPostRoutes(router *mux.Router) {
	router.HandleFunc("/recordSimcard", post.RecordSimcard).Methods("POST")
	router.HandleFunc("/recordStock", post.RecordStock).Methods("POST")
	router.HandleFunc("/postUser", post.PostUser).Methods("POST")
}
