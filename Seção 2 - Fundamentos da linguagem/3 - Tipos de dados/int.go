package main

import "fmt"

func main() {

	// Inteiros (int8, int16, int32, int64) <quantidade de bits que o int suporta>

	var numero int16 = 100
	fmt.Println(numero)

	//  Inteiro sem bits "int". Criar uma variável usando apenas o "int" ele automaticamente criará a variável de
	//  acordo com a arquitetura do computador que está executando, por exemplo, meu PC é 64 bits, logo ele criar uma variável int64...
	var numeroInteiro int = 100000
	// numeroInteiro := 10000 // Cria uma variável do tipo "int"
	fmt.Println(numeroInteiro)

	// Cria um inteiro sem aceitar valor negativo
	var numero2 uint32 = 1000
	fmt.Println(numero2)

	// alias, ou "apelido".

	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// INT8 = BYTE
	var numero4 byte = 123
	fmt.Println(numero4)

	// Inteiro ou float
	var intVazio int
	fmt.Println(intVazio) // Retorna 0
}
