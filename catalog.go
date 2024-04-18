package main

import "github.com/froehlichluana/loja-digport-backend/model"


func createCatalog() []model.Product {
	products := []model.Product{
		{
			Name: "Test",
			Description: "This is a test product",
			Category: "Test category",
			Id: "7493jr03",
			Price: 74.97,
			Quantity: 150,
			Image: "Test.jpg",
		},
		{
			Name: "Test 2",
			Description: "Another one test",
			Category: "Test 2 category",
			Id: "748jkf903",
			Price: 13.87,
			Quantity: 340,
			Image: "Test2.jpg",
		},
	}
	return products
}