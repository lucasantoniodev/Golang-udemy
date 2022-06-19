package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
	
)

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()

	// erro := aplicacao.Run(os.Args)
	// if erro != nil {
	// 	log.Fatal(erro)
	// }

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro, "OPA, TIVEMOS UM PROBLEMA!")
	}


}






// Comando da aplicação: 
					//< Comando > < --parâmetro > < valor do parâmetro >
	// go run main.go ip --host amazon.com.br
	
		// Para executar com o valor padrão usa o comando:
			// go run main.go ip


	// go run main.go servidores --host amazon.com.br