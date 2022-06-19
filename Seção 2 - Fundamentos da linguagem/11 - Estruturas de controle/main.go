package main

import "fmt"

func main() {
	fmt.Println("==========Estrutura de controle==========")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// IF INIT
	// Declarando uma variável e depois avaliando.
	// Esta variável está limitada ao escopo do IF
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que 0")
	} else if numero < -10 {
		fmt.Println("É menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}

	

}
