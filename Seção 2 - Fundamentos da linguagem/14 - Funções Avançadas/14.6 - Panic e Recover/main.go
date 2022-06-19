package main

import "fmt"

func recuperar() {

	if rec := recover(); rec != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}

}


func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperar()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("Média é 6!!!!!")
}

func main() { fmt.Println(alunoEstaAprovado(7, 6))
	println("Olá")
}
