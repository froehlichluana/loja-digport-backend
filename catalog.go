package main

import "github.com/froehlichluana/loja-digport-backend/model"

var Products []model.Product = []model.Product {}

func productsByName (name string) []model.Product { 
	result := []model.Product{}

	for _, product := range Products {
		if product.Name == name {
			result = append(result, product)
		}
	}
	return result
}

func registerProduct(product model.Product) {
	Products = append(Products, product)
}




