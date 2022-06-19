package enderecos

import "strings"

// TipoEndereço verifica se tem um endereço válido
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Estrada", "Rodovia"}

	primeiraPalavraDoEndereco := strings.Split(endereco, " ")[0]
	// Rua dos bobos
	// Slipt: ["Rua", "dos", "bobos"]

	enderecoTemUmTipoValido := false

	for _, value := range tiposValidos {
		if strings.EqualFold(value, primeiraPalavraDoEndereco) {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return primeiraPalavraDoEndereco
	}

	return "Tipo Inválido"
}
