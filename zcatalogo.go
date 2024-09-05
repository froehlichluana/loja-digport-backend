package main

import (
	"errors"

	"github.com/froehlichluana/loja-digport-backend/model"
)

var Products []model.Produto = []model.Produto {}

func productsByName (nome string) []model.Produto { 
	result := []model.Produto{}

	for _, product := range Products {
		if product.Nome == nome {
			result = append(result, product)
		}
	}
	return result
}

func registerProduct(newProduct model.Produto) error{
	for _, product := range Products {
		if newProduct.ID == product.ID{
			return errors.New("ID já cadastrado")
		}
	}

	Products = append(Products, newProduct)
	return nil
}




