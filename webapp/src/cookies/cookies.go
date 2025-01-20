package cookies

import (
	"net/http"
	"time"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Configurar utiliza as variáveis de ambiente para a criação do SecureCookie
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Salvar registra as informações de autenticação
func Salvar(w http.ResponseWriter, ID, token string) error {

	dados := map[string]string{
		"id":    ID,
		"token": token,
	}

	// codificando os dados para salvar no cookie
	dadosCodificados, erro := s.Encode("dados", dados)
	if erro != nil {
		return erro
	}

	// enviando os dados
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

// Ler retorna os valores armazenados no cookie
func Ler(r *http.Request) (map[string]string, error) {
	// pegando os cookies do navegador
	cookie, erro := r.Cookie("dados")
	if erro != nil {
		return nil, erro
	}

	// variável para receber os valores dentro do cookie
	valores := make(map[string]string)

	// descodificando o cookie
	if erro = s.Decode("dados", cookie.Value, &valores); erro != nil {
		return nil, erro
	}

	return valores, nil

}

// deletar remove os valores armazenados no cookie
func Deletar(w http.ResponseWriter) {
	// enviando os dados co cookie em branco
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
