package seguranca

import (
	"golang.org/x/crypto/bcrypt"
)

// Diferente de uma criptografia, o hash não tem como voltar "descriptografar", a unica maneira
// de indetificar é por comparação

func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha compara uma senha e um hash e retorna se elas são iguais
func VerificarSenha(senhaComHash, senhaString string) error {

	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}
