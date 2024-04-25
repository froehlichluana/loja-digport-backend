package main

import (
	"encoding/json"
	"net/http"

	"github.com/froehlichluana/loja-digport-backend/model"
)

func StartServer() {
	http.HandleFunc("/produtos", productsHandler)

	http.ListenAndServe(":8080", nil)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getProducts(w, r)
	}else if r.Method == "POST" {
		addProduct(w, r)
	}
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	json.NewDecoder(r.Body).Decode(&product)

	registerProduct(product)

	w.WriteHeader(http.StatusCreated)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	queryName := r.URL.Query().Get("name")
	if queryName != "" {
		productsFilteredByName := productsByName(queryName)
		json.NewEncoder(w).Encode(productsFilteredByName)
	}else{
		products := Products
		json.NewEncoder(w).Encode(products)
	}
}