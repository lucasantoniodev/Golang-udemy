package main

import "fmt"

func main() {
	// Arrays internos
	// make() é utilizado quando queremos alocar um espaço na memoria

	slice := make([]float32, 10, 11) // Tamanho de 10 e capacidade de 15
	fmt.Println(slice)

	slice = append(slice, 5)
	slice = append(slice, 6)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	fmt.Println("==================")

	slice1 := make([]float32, 5)
	fmt.Println(slice1)
	slice1 = append(slice1, 10)
	fmt.Println(slice1)

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}
