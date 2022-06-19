package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// o método é uma Função, porém ele está "grudado a uma estrutura" "(u usuario)"
func (u *usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Teco", 20}
	usuario1.salvar()

	usuario2 := usuario{"Tico", 21}
	usuario2.salvar()
	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
