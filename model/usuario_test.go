package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidaUsuario(t *testing.T) {
	t.Run("Deve retornar nulo quando usuário estiver completo", func(t *testing.T) {
		//arrange
		usuario := Usuario{Nome: "Jim", Email: "Jim@test.com", Senha: "73h93p0934"}

		//act
		err := usuario.Validar(usuario)

		//assert
		assert.NoError(t, err)

	})

	t.Run("Deve retornar erro quando usuário tiver senha vazia", func(t *testing.T) {
		//arrange
		usuario := Usuario{Nome: "Jim", Email: "jim@test.com", Senha: ""}
		expectedErrorMessage := "senha não pode ser vazia."
		//act
		err := usuario.Validar(usuario)

		//assert
		assert.Error(t, err)
		assert.EqualError(t, err, expectedErrorMessage)

	})

	t.Run("Deve retornar erro quando usuário tiver email vazia", func(t *testing.T) {
		//arrange
		usuario := Usuario{Nome: "Jim", Email: "", Senha: "0384r93j"}
		expectedErrorMessage := "email do usuário não pode ser vazio"
		//act
		err := usuario.Validar(usuario)

		//assert
		assert.Error(t, err)
		assert.EqualError(t, err, expectedErrorMessage)

	})

	t.Run("Deve retornar erro quando usuário tiver nome", func(t *testing.T) {
		//arrange
		usuario := Usuario{Nome: "", Email: "jim@test.com", Senha: "0384r93j"}
		expectedErrorMessage := "nome do usuário não pode ser vazio"
		//act
		err := usuario.Validar(usuario)

		//assert
		assert.Error(t, err)
		assert.EqualError(t, err, expectedErrorMessage)

	})

}


