package main

import "fmt"

func clousure() func() {

	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaNova := clousure()
	funcaNova()
}