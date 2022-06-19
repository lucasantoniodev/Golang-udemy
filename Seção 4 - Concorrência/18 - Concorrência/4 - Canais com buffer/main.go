package main

import "fmt"

func main() {
	// Quando não usa buffer, ocorre o deadlock pois ele atua como bloqueante
	canal := make(chan string, 2) // buffer de 2, ou seja, só vai bloquear quando atingir a capacidade máxima

	canal <- "Olá mundo"
	canal <- "Hello World"

	mensagem := <- canal
	mensagem2 := <- canal
	
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}