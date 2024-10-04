package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/froehlichluana/loja-digport-backend/model"
)

var jwtKey = []byte("secret")

func CriaUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	var usuario model.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)

	err := usuario.Validar()
	if err != nil {
		fmt.Println("Usuário informado inválido:", err)
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(model.Erro{Mensagem: err.Error()})
		return
	}

	err = model.CriaUsuario(usuario)
	if err != nil {
		fmt.Println("Erro ao criar usuário: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		fmt.Println("Usuário criado")
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func BuscaUsuarioPorEmail(w http.ResponseWriter, r *http.Request){

		email := r.URL.Query().Get("email")

		usuario, err := model.BuscaUsuarioPorEmail(email)
		if err != nil {
			fmt.Println("Erro ao buscar usuário: ", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(usuario)

	}

	func LoginHandler(w http.ResponseWriter, r *http.Request){
		var usuario model.Usuario
		json.NewDecoder(r.Body).Decode(&usuario)
	
		username := usuario.Email
		senhatxt := usuario.Senha
		
		user, error := model.BuscaUsuarioPorEmail(username)
		if error != nil{
			w.WriteHeader(http.StatusNotFound)
			return
		}

		hash := user.Senha

		error = model.ValidaLogin(hash, senhatxt)

		if error != nil {
			//fmt.Println("Erro ao validar senha:")
			//w.WriteHeader(http.StatusNotFound)
			dataExpiracao := time.Now().Add(5 * time.Minute)
		standardToken := &jwt.StandardClaims{
			ExpiresAt: dataExpiracao.Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, standardToken)
		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			fmt.Println("Erro ao validar jwt:")
			w.WriteHeader(http.StatusUnauthorized)

			return 

		}

		w.Write([]byte(tokenString))
		}else{
			w.WriteHeader(http.StatusUnauthorized)
		}

		

	}

	

	func AtualizaUsuario(w http.ResponseWriter, r *http.Request){
		
	}

