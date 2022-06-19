package main

import "fmt"

// UMA STRUCT É MUITO SIMILAR A UMA CLASSE

type usuario struct {
	nome     string
	idade    uint8
	endereco address
}

type address struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	// MÉTODO 1
	var u usuario
	fmt.Println(u) // Resultado vazio string = "", uint8 "0"

	u.nome = "Teco"
	u.idade = 22
	fmt.Println(u)
	// FIM DO MÉTODO 1

	// MÉTODO 2 MAIS RÁPIDO
	var endereco address = address{"Rua dos bobos", 5}

	user2 := usuario{"Tico", 21, endereco}
	fmt.Println(user2)
	// FIM DO MÉTODO 2

	// MÉTODO 3 É QUANDO NÃO TEMOS TODOS OS DADOS
	user3 := usuario{nome: "Lucas"}
	fmt.Println(user3)

}
