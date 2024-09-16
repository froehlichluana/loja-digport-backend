package model

import (
	"fmt"

	"github.com/froehlichluana/loja-digport-backend/db"
	"golang.org/x/crypto/bcrypt"
)

type Usuario struct{
	ID string `json:"id"`
	Nome string	`json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
	Telefone string `json:"telefone"`
	Endereco string `json:"endereco`

}

func (u Usuario) Validar () error {
	if u.Nome == "" {
		return fmt.Errorf("nome não pode ser vazio")
	}
	if u.Email == "" {
		return fmt.Errorf("email não pode ser vazio")
	}
	if u.Senha == "" {
		return fmt.Errorf("senha não pode ser vazia")
	}
	return nil
}

func CriaUsuario (usuario Usuario) error{
	hash, err := hashPassword(usuario.Senha)
	if err != nil {
		return fmt.Errorf("Erro ao criar usuário: %w", err)
	}

	db := db.ConectaBancoDados()
	defer db.Close()

	_, err = db.Exec("INSERT INTO USUARIO  VALUES ($1, $2, $3, $4, $5)",usuario.Nome, usuario.Email, hash, usuario.Telefone, usuario.Endereco)
	if err != nil {
		fmt.Errorf("erro ao inserir usuário no banco de dados: %w", err)
	}
	return nil
}

func hashPassword(senha string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), 14)
	if err != nil {
		return "", fmt.Errorf("erro ao tentar gerar hash da senha: %w", err)
	}
	return string(bytes), nil
}

