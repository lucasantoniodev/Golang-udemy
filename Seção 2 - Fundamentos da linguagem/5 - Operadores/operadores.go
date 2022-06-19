package main

import "fmt"

func main() {
	// Operadores aritimédicos: + - / * %
	soma := 1 + 2
	subtracao := 5 - 3
	divisao := 10 / 3
	multiplicacao := 10 * 5
	restodaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restodaDivisao)

	// Sobre comparação de dados ou soma de variáveis
	// var numero1 int32 = 10
	var numero1 int64 = 10 // Agora sim é permitido pois os tipos são iguais
	var numero2 int64 = 25
	soma2 := numero1 + numero2
	println(soma2)
	// FIM OPERADORES ARITIMÉTICOS

	fmt.Println("=====================")

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	fmt.Println(1 == 2)
	// FIM OPERADORES RELACIONAIS

	fmt.Println("=====================")

	// OPERADORES LÓGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	// FIM OPERADORES LÓGICOS

	fmt.Println("=====================")

	// OPERADORES UNÁRIOS
	numero := 10
	fmt.Println(numero)

	numero++
	fmt.Println(numero)
	numero += 1
	fmt.Println(numero)

	numero--
	fmt.Println(numero)
	numero -= 1
	fmt.Println(numero)

	numero *= 3
	fmt.Println(numero)

	numero /= 10
	fmt.Println(numero)

	numero %= 3
	fmt.Println(numero)
	// FIM OPERADORES UNÁRIOS

	fmt.Println("=====================")

	// OPERADORES TERNÁRIOS NÃO TEM NO GOLANG POIS A SUA PREMISSA DIZ QUE "SÓ A UM JEITO DE FAZER UMA COISA, TRAZER O SIMPLES"

	// NÃO EXISTE: text := (5 > 4) ? "MAIOR" : "NÃO É MAIOR"

	// SOLUÇÃO:
	var text string

	if 2 > 5 {
		text = "Maior"
	} else {
		text = "Menor"
	}

	fmt.Println(text)

}
