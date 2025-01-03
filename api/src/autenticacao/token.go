package autenticacao

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken gera um token assinado e com as permissões do usuário
func CriarToken(usuarioID uint64) (string, error) {

	// definindo as permissões do token
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID

	// gerando uma chave para o token poder ser validado - secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	return token.SignedString([]byte("Secret")) // gerando o secret e retornando
}
