package main

import "fmt"

func main() {

	// Função anônima
	func() {
		fmt.Println("Olá mundo!")
	}()

	// Passando parâmetros em funções anônimas
	func(value string) {
		fmt.Println(value)
	}("Olá")

	// Retornando um texto formatado com variável
	retorno := func(text string) string {
		return fmt.Sprintf("Recebido => %s", text)
	}("Passando parâmetro")

	fmt.Println(retorno)
}
