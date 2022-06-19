package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nomes de servidor na internet"
	
	flags := []cli.Flag{ // Flags são parâmetros do comando
		cli.StringFlag{
			Name:  "host", // Nome do parâmetro
			Value: "devbook.com.br", // Valor padrão do comando
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ip", // Nome do comando
			Usage: "Busca IPS de endereços na internet", // Descrição
			Flags: flags,
			Action: buscarIps,
		},
		{
			Name: "servidores",
			Usage: "Busca o nome dos servidores na internet",
			Flags: flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host") // Capturando o valor da flag "host"

	// pacote para buscar ips
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro, "\nOPA, TIVEMOS UM PROBLEMA!")
	}
	
	// Podemos iterar sobre a variável "ips" pois ele é um slice, sendo possível um host retornar vários IPs
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host") // Capturando o valor da flag "host"

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}

}