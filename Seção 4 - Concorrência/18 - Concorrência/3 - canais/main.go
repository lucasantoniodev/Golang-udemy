package main

import (
	"fmt"
	"time"
)

func main() {
	// Canal de comunicação, pode enviar e receber dados
	canal := make(chan string) // só vai trafegar dados do tipo "string"
	go escrever("Olá mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	// SIMPLIFICANDO

	// for {
	// 	mensagem, aberto := <- canal // O canal espera receber um valor
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // Enviando um valor para o canal
		time.Sleep(time.Second)
	}
	close(canal)
}

// deadlock é quando você não tem mais nenhum lugar que está envinando dado para o canal, só que o canal ainda espera receber um dado, isso é um PROBLEMA!
// ELE NÃO DETECTÁVEL EM COMPILAÇÃO, SÓ EM EXECUÇÃO
