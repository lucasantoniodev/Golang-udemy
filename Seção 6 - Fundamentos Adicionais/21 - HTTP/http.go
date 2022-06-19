package main

import (
	"log"
	"net/http"
)


func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func main() {
	// HTTP é um protocolo de comunicação - Base da comunicação web

	// Client (Request) - Servidor (Response)

	// Rotas (Identifica os tipos de mensagens que está sendo enviada e o tipo de processamento)
	// URI - Identificador do recurso
	// Método - GET, POST, PUT, DELETE

	http.HandleFunc("/home", home)

	
	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregar página de usuários"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))


}