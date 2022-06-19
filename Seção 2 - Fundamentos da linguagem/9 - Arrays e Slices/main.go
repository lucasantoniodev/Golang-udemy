package main

import (
	"fmt"
	// "reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	// ARRAYS
	// OBRIGATÓRIO ESPECIFICAR O TAMANHO PARA ARRAY
	var array1 [5]int
	fmt.Println(array1)

	array1[0] = 10
	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	var array3 [5]string = [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array3)

	// Array semi-flexível, ou seja, ele vai fixar o tamanho de acordo com a quantidade de posições que foi inicializado
	array4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array4)
	// FIM ARRAYS

	fmt.Println("===================")

	// Slices
	// Slice não é array, parece, mas não é, ele aponta pra um array, como se fosse um ponteiro apontando para um array que ele referência "uma fatia de array"
	slice := []int{2, 3, 4, 5, 6}
	fmt.Println(slice)

	// fmt.Println(reflect.TypeOf(slice))
	// fmt.Println(reflect.TypeOf(array3))

	// Adicionando elemento e retornando um slice novo
	slice = append(slice, 18)
	fmt.Println(slice)

	// Criando um slice com base em um array
	slice2 := array3[1:3]
	fmt.Println(slice2)

	array3[1] = "Posição alterada" // Alterando valor da array
	fmt.Println(slice2) // Porém o slice já consegue capturar esse valor
	// ISSO SÓ É POSSÍVEL PORQUE O SLICE É UMA "REFERÊNCIA/PONTEIRO" DE UMA ARRAY

}
