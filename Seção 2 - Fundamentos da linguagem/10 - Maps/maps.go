package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// Todas as chaves precisam ser do mesmo tipo e todos os valores precisam ser de tipos iguais

	usuario := map[string]string{
		"nome":      "Teco",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{

		"nome": {
			"primeiro": "Lucas",
			"último":   "Conceição",
		},
		"curso": {
			"nome":   "engenharia",
			"campus": "uninter",
		},
	}

	fmt.Println(usuario2)

	// Remover chave
	delete(usuario2, "curso")
	fmt.Println(usuario2)

	// Adicionando uma nova chave
	usuario2["age"] = map[string]string{
		"age": "22",
	}
	fmt.Println(usuario2)
}
