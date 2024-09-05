package main

import ("testing"
		"github.com/stretchr/testify/assert"
		"github.com/stretchr/testify/require"

)

func TestSoma(t *testing.T){
	t.Run("Should sum two numbers.", func(t *testing.T) {
		//arrange
		num1 := 4
		num2 := 8

		expectedResult := 12

		// act
		result := soma(num1, num2)

		//assert
		assert.Equal(t, expectedResult, result)

	})

}

func TestValidaNome(t *testing.T){

	t.Run("Should return an error when name is empty", func(t *testing.T) {
		//arrange <=> given
		name := ""

		expectedErrorMessage := "Nome não preenchido"

		//act <=> when
		err := validaNome(name)

		//assert <=> then
		require.Error(t, err)
		assert.Equal(t, expectedErrorMessage, err.Error())

	})

	t.Run("Should not return an error when name is filled", func(t *testing.T) {
		//arrange
		name := "produto"

		//act
		err := validaNome(name)

		//assert
		assert.Nil(t, err)

	})
}