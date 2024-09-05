package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/froehlichluana/loja-digport-backend/controller"
	"github.com/rs/cors"
	
)

func HandleRequests() {
	route := mux.NewRouter()

	route.HandleFunc("/produto", controller.BuscaProdutosHandler).Methods("GET")
	route.HandleFunc("/produto", controller.BuscaProdutoPorNomeHandler).Methods("GET")
	route.HandleFunc("/produto", controller.CriaProdutosHandler).Methods("POST")
	route.HandleFunc("/produto/{id}", controller.RemoveProdutoHandler).Methods("DELETE")
	route.HandleFunc("/produto", controller.AtualizaProdutoHandler).Methods("PUT")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
		AllowedHeaders: []string{"Content-Type"},
		AllowCredentials: true,
	})
	
	handler := c.Handler(route)
	

	http.ListenAndServe(":8080", handler)
}

