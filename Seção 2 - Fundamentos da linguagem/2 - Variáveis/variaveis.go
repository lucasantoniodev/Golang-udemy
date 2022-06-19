package main

import (
	"fmt"
	"strings"
)

func main() {
	// Toda variável e pacote criado/importado e não for usado gera um erro e impede a execução

	// "!" Variável criado somente var nome = "Teco" é permitido para criar fora de funções, tornando ela do escopo global diferente do := que é pra escopo local

	// Criando variável com tipo explícito
	var variavel1 string = "Variável 1"
	fmt.Println(variavel1)

	// Criando variável com tipo implícito
	// Nome técnico do ":=" é "Inferência de tipo"
	variavel2 := "Variável 2"
	fmt.Println(variavel2)

	// Criando um conjunto de variáveis com tipo explícito
	var (
		variavel3 string = "lala"
		variavel4 string = "lalala"
	)
	fmt.Println(variavel3, variavel4)

	// Criando um conjunto de variáveis com tipo implícito
	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	// Criando uma constante com tipo explícito
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	// Criando um conjunto de constante
	const (
		constante2 string = "Constante 2"
	)
	fmt.Println(constante2)

	// Invertendo valor das variáveis
	variavel5, variavel2 = variavel2, variavel5
	println(variavel2, variavel5)

	testando1 := "Oi"
	testando2 := "oi"

	if strings.EqualFold(testando1,testando2) {
		fmt.Println("OLÁ MEU CHAPA")
	} else {
		fmt.Println("NÃO DEU MEU CHAPA")
	}
	
}