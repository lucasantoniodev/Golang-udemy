package main

import (
	"errors"
	"fmt"
)

func main() {
	// Tipo de erro

	var erro error
	fmt.Println(erro) // Retorna <nil> = nulo

	// PACOTE NATIVO DO GO PARA TRATAMENTO DE ERROS
	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)
}