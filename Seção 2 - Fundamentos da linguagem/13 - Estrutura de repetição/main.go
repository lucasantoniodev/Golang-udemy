package main

import (
	"fmt"
	
)

func main() {
	fmt.Println("================ for ================")

	var i int = 0

	for i < 10 {
		i++

		fmt.Println(i)
	}
	fmt.Println(i)

	fmt.Println("================")

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando: ", j)
		// time.Sleep(time.Second)
	}

	fmt.Println("================")

	// Iterando sobre Array
	nomes := [3]string{"Teco", "Tico", "Esquilo"}
	for index, value := range nomes {
		fmt.Println(index, value)
	}

	fmt.Println("================")

	// Iterando sobre string
	for indice, letra := range "Palavra" {
		fmt.Println(indice, string(letra))
	}

	// Iterando sobre mapas
	usuario := map[string]string {
		"nome": "Lucas",
		"sobrenome": "Silva",
	}

	for key, value := range usuario {
		fmt.Println(key,value)
	}


}
