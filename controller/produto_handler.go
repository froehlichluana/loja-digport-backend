package controller

import (
	"encoding/json"
	"net/http"

	"github.com/froehlichluana/loja-digport-backend/model"
	"github.com/gorilla/mux"
)

func BuscaProdutosHandler(w http.ResponseWriter, r *http.Request) {
	produtos := model.BuscaTodosProdutos()
	json.NewEncoder(w).Encode(produtos)
}

func BuscaProdutoPorNomeHandler(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")
	produto := model.BuscaProdutoPorNome(nome)
	json.NewEncoder(w).Encode(produto)

}

func CriaProdutosHandler(w http.ResponseWriter, r *http.Request) {
	var Produto model.Produto
	json.NewDecoder(r.Body).Decode(&Produto)


	error := model.CriaProduto(Produto)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
	}else{
		w.WriteHeader(http.StatusCreated)
	}
	
}

func RemoveProdutoHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	err := model.RemoveProduto(id)
	if err != nil {
		userError := model.Erro{Mensagem: "ocorreu um erro ao tentar excluir o produto"}
		json.NewEncoder(w).Encode(userError)
		w.WriteHeader(http.StatusInternalServerError)
		return 
	} 
	w.WriteHeader(http.StatusNoContent)
}

func AtualizaProdutoHandler(w http.ResponseWriter, r *http.Request){
	
}
 