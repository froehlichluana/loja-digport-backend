package main

import (
	"encoding/json"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/produtos", produtosHandler)

	http.ListenAndServe(":8080", nil)
}

func produtosHandler (w http.ResponseWriter, r *http.Request) {
	produtos := createCatalog()

	json.NewEncoder(w).Encode(produtos)
}