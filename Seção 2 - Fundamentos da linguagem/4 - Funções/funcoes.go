package main

import "fmt"

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	// Executando função com outra função dentro
	anyFunction()

	// Recebendo dois retornos de vez de uma função
	// Importante seguir a ordem dos retornos
	resultado1, resultado2 := calculosMatematicos(10, 15)
	fmt.Println(resultado1, resultado2)

	// Omitindo valores quando uma função retorna mais de um valor mas não vamos utilizar
	resultado3, _ := calculosMatematicos(10, 15)
	fmt.Println(resultado3)
}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func anyFunction() {
	var f = func(txt string) string {
		fmt.Println(txt)
		return "Função concluída"
	}

	status := f("Imprimindo texto")
	fmt.Println("Status: ", status)

}

// No go é possível retornar mais de um retorno por funcao
// Todos os parâmetros pega o último parâmetro, caso vc não expecifíque individualmente
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}
