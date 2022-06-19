package main

import "fmt"

func diaDaSemana(numero int) string {

	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Dia inválido"
	}
}

func mes(numero int) string {

	switch {
	case numero == 1:
		return "Janeiro"
	case numero == 2:
		return "Fevereiro"
	case numero == 3:
		return "Março"
	case numero == 4:
		return "Abril"
	case numero == 5:
		return "Maio"
	case numero == 6:
		return "Junho"
	case numero == 7:
		return "Julho"
	case numero == 8:
		return "Agosto"
	case numero == 9:
		return "Setembro"
	case numero == 10:
		return "Outubro"
	case numero == 11:
		return "Novembro"
	case numero == 12:
		return "Dezembro"
	default:
		return "Mês inválido"

	}

}

// Não precisamos usar o break no GO
func clausula(numero int) string {
	// CASO QUEIRA PASSAR PARA UMA OUTRA CASE USE "FALLTROUGTH"

	var armazenaValor string

	switch numero {
	case 1:
		armazenaValor = "Valor 1"
		fallthrough

	case 2:
		armazenaValor = "Valor 2"

	case 3:
		armazenaValor = "Valor 2"

	}

	return armazenaValor
}

func main() {
	fmt.Println("====== Switch ======")
	var dia string = diaDaSemana(2)
	var mes string = mes(5)

	fmt.Println(dia, mes)
	fmt.Println(clausula(1))
}
