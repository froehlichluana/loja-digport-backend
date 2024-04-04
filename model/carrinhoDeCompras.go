package model

type Cart struct{
	Id string
	UserId string
	Total float64
	ProductInfo []ProductInfo
	Quantity int
}

type ProductInfo struct{
	ProductId string
	ProductQuantity int
}