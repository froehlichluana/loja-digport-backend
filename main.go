package main

import (
	"fmt"
)


func main() {
	StartServer()

	fmt.Println("Hello, world!")
	
	fmt.Printf("Esse é o catálogo da loja Neon Bastet: +%v", createCatalog())
}
