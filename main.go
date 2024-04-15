package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"golang.org/x/time/rate"

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

	// Criar um limitador de taxa com 1000 requisições por segundo
	limiter := rate.NewLimiter(rate.Limit(1000), 1000)

	// Adicionar middleware para rate limiting
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !limiter.Allow() {
				http.Error(w, "Too many requests", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	log.Fatal(http.ListenAndServe(":60060", headers.AddCorsHeaders(router)))
}
