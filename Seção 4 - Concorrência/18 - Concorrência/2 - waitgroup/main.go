package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(4) // Colocando 2 goroutines na fila

	
	go func() {
		escrever("1")
		waitGroup.Done() // Tirar função do contador "-1"
	}()

	go func() {
		escrever("2")
		waitGroup.Done()
	}()

	go func() {
		escrever("3")
		waitGroup.Done()
	}()

	go func() {
		escrever("4")
		waitGroup.Done()
	}()
	
	waitGroup.Wait() // Esperar a contagem das goroutines chegar em 0 para finalizar

}

func escrever(texto string) {
	for i := 0; i < 2; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}
