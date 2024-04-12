package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"kdl_api_rest_internal/headers"
	routesDelete "kdl_api_rest_internal/routes/routes_delete"
	routesGet "kdl_api_rest_internal/routes/routes_get"
	routesPost "kdl_api_rest_internal/routes/routes_post"
)

func main() {
	router := mux.NewRouter()

	routesGet.SetGetRoutes(router)
	routesPost.SetPostRoutes(router)
	routesDelete.SetDeleteRoutes(router)

	log.Fatal(http.ListenAndServe(":60060", headers.AddCorsHeaders(router)))
}
