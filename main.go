package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"kdl_api_rest_internal/db"
	"kdl_api_rest_internal/headers"
	routesGet "kdl_api_rest_internal/routes/routes_get"
	routesPost "kdl_api_rest_internal/routes/routes_post"
)

func main() {
	db.ConnectMysql()
	router := mux.NewRouter()

	routesGet.SetGetRoutes(router)
	routesPost.SetPostRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", headers.AddCorsHeaders(router)))
}
