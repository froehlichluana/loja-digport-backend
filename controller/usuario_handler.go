package controller

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/froehlichluana/loja-digport-backend/model"
)

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
