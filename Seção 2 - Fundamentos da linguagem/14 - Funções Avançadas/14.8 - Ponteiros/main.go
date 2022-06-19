package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1

}

func inverterSinalComPonteiro(numero *int) {
	// Desreferenciando
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40

	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero) // Passando o endereço de memória onde está salva
	fmt.Println(novoNumero)
}
