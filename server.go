package main

import (
	"github.com/froehlichluana/loja-digport-backend/routes"
)

func StartServer() {
	routes.HandleRequests()
}