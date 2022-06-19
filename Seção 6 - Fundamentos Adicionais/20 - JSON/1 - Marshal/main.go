package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	// Converter Map/Struct para Json
	dog := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(dog)

	dogJSON, erro := json.Marshal(dog)

	if erro != nil {
		log.Fatal(erro)
	}

	// Retorna um array de bytes
	fmt.Println(dogJSON)

	// Transforma um array de bytes em formato esperado
	fmt.Println(bytes.NewBuffer(dogJSON))

	dog2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	dog2JSON, erro := json.Marshal(dog2)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(dog2JSON)
	fmt.Println(bytes.NewBuffer(dog2JSON))

}
