package main

import "fmt"

func main() {

	totalDaSoma := soma(1, 2, 3, 4, 5, 6)
	fmt.Println(totalDaSoma)

	escrever("Olá mundo", 1,2,3,4)

}

// Recebendo vários parâmetros, a variável se torna um slice, sendo possível iterar
func soma(numeros ...int) int {
	total := 0

	// iterando sobre o slice
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(text string, numeros ...int) {

	for _, value := range numeros {
		fmt.Println(text, value)

	}

}
