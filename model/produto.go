package model

type Product struct{
	Name string `json:"name"`
	Description string `json:"description"`
	Category string `json:"category"`
	Id string `json:"id"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
	Image string `json:"image"`
}