package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

// OMITIMOS O NOME DA VARIÁVEL POIS QUEREMOS APENAS UMA "CÓPIA"
// NÃO QUEREMOS ACESSAR "estudante.pessoa.nome" E SIM "estudante.nome"
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Lucas", "Antônio", 22, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "UNINTER"}
	fmt.Println(e1)
	fmt.Println(e1.nome)
	fmt.Println(e1.pessoa.nome)
}
