package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

	return token.SignedString([]byte(config.SecretKey)) // gerando o secret e retornando
}

// ValidarToken verifica se o token passado na requisição é valido
func ValidarToken(r *http.Request) error {

	// capturando o token
	tokenString := extrairToken(r)

	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}

	// validando o token
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

// ExtrairUsuarioId retorna o usuarioId que está salvo no token
func ExtrairUsuarioId(r *http.Request) (uint64, error) {
	// capturando o token
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return 0, erro
	}

	// extraindo as permissões
	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// convertendo valor para uint64
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return usuarioID, nil
	}

	return 0, errors.New("token inválido")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {

	// verificando se o método de assinatura do token é do tipo HMAC (que é um algoritmo de hash com chave compartilhada).
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método ded assinatura inesperado! %s", token.Header["alg"])
	}

	return config.SecretKey, nil
}
