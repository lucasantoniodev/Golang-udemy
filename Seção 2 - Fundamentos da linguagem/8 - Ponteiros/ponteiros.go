package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ponteiros")

	// APENAS ENTENDENDO COMO FUNCIONA UMA CÓPIA DE VALOR
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)
	variavel1++
	fmt.Println(variavel1, variavel2) // ALTERA APENAS A VARIÁVEL 1

	fmt.Println("=======================")

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int
	var ponteiro *int
	fmt.Println(variavel3, ponteiro)

	variavel3 = 100
	ponteiro = &variavel3 // CAPTURANDO REFERÊNCIA NA MEMÓRIA
	fmt.Println(variavel3, ponteiro)

	desreferenciando := *ponteiro
	fmt.Println(desreferenciando) // LENDO VALOR QUE ESTÁ NA MEMÓRIA

	// MESMO ALTERANDO O VALOR DA VARIÁVEL A REFERÊNCIA DE MEMORIA NÃO MUDA
	variavel3 = 150
	ponteiro = &variavel3
	desreferenciando = *ponteiro
	fmt.Println(desreferenciando, ponteiro)

	// PONTEIRO NÃO GUARDA VALOR, ELE GUARDA UMA REFÊNCIA DE MEMÓRIA

	var aprendendo int = 10
	var aponteiro *int = &aprendendo

	fmt.Println(aponteiro) // Rerência na memória

}
