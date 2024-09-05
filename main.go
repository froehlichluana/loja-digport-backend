package main

import (
	"github.com/froehlichluana/loja-digport-backend/db"
)

func main() {

	db.InitDB()
	StartServer()

}
