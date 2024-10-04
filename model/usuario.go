package model

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/froehlichluana/loja-digport-backend/db"
	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	ID       string `json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Senha    string `json:"senha"`
	Telefone string `json:"telefone"`
	Endereco string `json:"endereco"`
}


func (u Usuario) Validar() error {
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

func CriaUsuario(usuario Usuario) error {
	hash, err := hashPassword(usuario.Senha)
	if err != nil {
		return fmt.Errorf("erro ao criar usuário: %w", err)
	}

	db := db.ConectaBancoDados()
	defer db.Close()

	_, err = db.Exec("INSERT INTO USUARIO (nome, email, senha, telefone, endereco) VALUES ($1, $2, $3, $4, $5)", usuario.Nome, usuario.Email, hash, usuario.Telefone, usuario.Endereco)
	if err != nil {
		//fmt.Errorf("erro ao inserir usuário no banco de dados: %w", err)
		log.Fatal(err)

	}
	return nil
}

func hashPassword(senha string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), 14)
	if err != nil {
		return "", fmt.Errorf("erro ao tentar gerar hash da senha: %w", err)
	}
	return string(bytes), nil
}

func ValidaLogin (hash string, senhatxt string) error {
	err:=bcrypt.CompareHashAndPassword([]byte(hash), []byte(senhatxt))

	if err != nil{
		return fmt.Errorf("senha inválida. usuário não autorizado")
	}
	return nil
}

func BuscaUsuarioPorEmail(email string) (*Usuario, error) {
	db := db.ConectaBancoDados()
	defer db.Close()

	var usuario Usuario
	err := db.QueryRow("SELECT id, nome, senha, email telefone, endereco FROM usuario WHERE email = $1", email).Scan(&usuario.ID, &usuario.Nome, &usuario.Senha, &usuario.Telefone, &usuario.Endereco)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario não encontrado %s", email)
		}
		return nil, err
	}
	return &usuario, nil
}


func AtualizaUsuario(usuario Usuario) error {
    db := db.ConectaBancoDados()
    defer db.Close()
 
    id := usuario.ID
    nome := usuario.Nome
    email := usuario.Email
    //senha := usuario.Senha
    telefone := usuario.Telefone
    endereco := usuario.Endereco
 
    result, err := db.Exec("UPDATE produtos SET nome= $1, email= $2, telefone= $3, endereco= $4 where id= $5", nome, email, telefone, endereco)
    if err != nil {
        panic(err.Error())
    }
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        panic(err.Error())
    }
 
    if rowsAffected == 0 {
        return fmt.Errorf("produto não encontrado")
    }
 
    fmt.Printf("Produto %s atualizado com sucesso (%d row affected)\n", id, rowsAffected)
 
    return nil
}
