package main

import (
	"fmt"
	"time"
)

func main() {
	

	go escrever("Olá mundo!") // usando goroutine
	escrever("Programando em Go!")

}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}

// CONCORRÊNCIA != PARALELISMO

	// CONCORRENTE: É quando os processos/tarefas revezam tempo para o uso do processador
	// PARALELISMO: É quando duas ou mais tarefas estão sendo executadas ao mesmo tempo


// Quando usamos o goroutine com o "go" na frente da função, estamos dizendo ao programa que execute a próxima função independente da "go" função terminar ou não