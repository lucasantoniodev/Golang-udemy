package enderecos_test

import (
	. "introducao-testes/enderecos" // COLOCAR O . na frente já trás com apelido, então o go entende que não precisa colocar o "pacote.NOMEDAFUNCAO"
	"strings"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) { // Test<NOMEDAFUNÇÃO>(t *testing.T)
	t.Parallel()

	cenariosDeTestes := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida ABC", "Avenida"},
		{"Rodovia ABC", "Rodovia"},
		{"Praça ABC", "Tipo Inválido"},
		{"Estrada ABC", "Estrada"},
		{"RUA ABC", "Rua"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTestes {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if !strings.EqualFold(retornoRecebido, cenario.retornoEsperado) {
			t.Errorf("O tipo de recebido %s é diferente do esperado %s",
				retornoRecebido,
				cenario.retornoEsperado,
			)
		}

	}

	// enderecoParaTeste := "avenida Paulista"
	// TipoDeEnderecoEsperado := "Avenida"
	// TipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)
	// if !strings.EqualFold(TipoDeEnderecoRecebido, TipoDeEnderecoEsperado) {
	// 	t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
	// 		TipoDeEnderecoEsperado,
	// 		TipoDeEnderecoRecebido,
	// 	)
	// }

}

func TestQualquer(t *testing.T) {
	t.Parallel() // Executar em paralelo
	if 1 > 2 {
		t.Errorf("Teste quebrou")
	}
}


// TESTE DE UNIDADE
// Menor unidade do seu código, neste caso uma função
