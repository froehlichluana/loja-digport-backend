package main

import "github.com/froehlichluana/loja-digport-backend/model"


func createCatalog() []model.Product {
	products := []model.Product{
		{
			Name: "Cama Janeleira",
			Description: "Cama janeleira para gatos em mdf",
			Category: "Caminhas",
			Id: "7493jr03",
			Price: 74.97,
			Quantity: 150,
			Image: "CamaJaneleira.jpg",
		},
		{
			Name: "Cama com Ventosa",
			Description: "Cama suspensa para janela, vidro e azulejo com ventosas",
			Category: "Caminhas",
			Id: "748jkf90",
			Price: 236.87,
			Quantity: 340,
			Image: "CamaVentosa.jpg",
		},
		{
			Name: "Cama Rede",
			Description: "Cama alta rede em madeira lavável",
			Category: "Caminhas",
			Id: "j3993dh7",
			Price: 269.99,
			Quantity: 107,
			Image: "CamaAlta.jpg",
		},
		{
			Name: "Cama Nuvem",
			Description: "Cama pelúcia redonda",
			Category: "Caminhas",
			Id: "wk978557",
			Price: 433.62,
			Quantity: 97,
			Image: "CamaNuvem.jpg",
		},
		{
			Name: "Fonte Automática",
			Description: "Fonte de água elétrica automática",
			Category: "Comedoures e bebedoures",
			Id: "39f38736",
			Price: 299.70,
			Quantity: 178,
			Image: "BebedouroAutomatico.jpg",
		},
		{
			Name: "Bebedouro Porcelana",
			Description: "Bebedouro porcelana para gatos",
			Category: "Comedoures e bebedoures",
			Id: "yh7520428",
			Price: 67.99,
			Quantity: 270,
			Image: "BebedorPorcelana.jpg",
		},
		{
			Name: "Comedouro Porcelana",
			Description: "Comedouro porcelana concavo",
			Category: "Comedoures e bebedoures",
			Id: "wpo10374",
			Price: 76.87,
			Quantity: 340,
			Image: "ComedouroPorcelana.jpg",	
		},
		{
			Name: "Comedouro Alto",
			Description: "Comedouro alto em madeira e porcelana",
			Category: "Comedoures e bebedoures",
			Id: "jdow9853",
			Price: 124.67,
			Quantity: 100,
			Image: "ComedouroAlto.jpg",
		},
		{
			Name: "Catnip spray",
			Description: "Spray de erva do gato/catnip",
			Category: "Brinquedos",
			Id: "dgw34509",
			Price: 27.56,
			Quantity: 540,
			Image: "CatnipSpray.jpg",
		},
		{
			Name: "Brinquedo Rato",
			Description: "Brinquedo rato com catnip",
			Category: "Brinquedos",
			Id: "4685kr78",
			Price: 31.45,
			Quantity: 430,
			Image: "RatoCatnip.jpg",
		},
		{
			Name: "Brinquedo fita",
			Description: "Brinquedo varinha com fita cetim para gatos",
			Category: "Brinquedos",
			Id: "ejdw3456",
			Price: 46.99,
			Quantity: 178,
			Image: "BrinquedoFita.jpg",
		},
		{
			Name: "Arranhador Papelão",
			Description: "Arranhador de papelão para gatos em formato rampa",
			Category: "Brinquedos",
			Id: "typ71290",
			Price: 87.69,
			Quantity: 87,
			Image: "Arranhador.jpg",
		},

	}
	return products

}

