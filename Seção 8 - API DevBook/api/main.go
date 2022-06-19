package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	r := router.GerarRotas()

	porta := fmt.Sprintf(":%d", config.Porta)
	fmt.Printf("Escutando na porta %s", porta)
	log.Fatal(http.ListenAndServe(porta, r))
}
