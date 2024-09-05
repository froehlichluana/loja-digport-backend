package main

import "fmt"

func soma (i, j int) int{
	return i + j
}

func validaNome(nome string) error{

	if nome == ""{
	return fmt.Errorf("Nome não preenchido")
}
	return nil

}