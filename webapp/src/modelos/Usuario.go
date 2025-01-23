package modelos

import (
	"net/http"
	"time"
)

// Usuário representa uma pessoa utilizando a rede social
type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	CriadoEm    time.Time    `json:"criadoEm"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}

// BuscarUsuarioCompleto faz 4 requisições na api para montar o usuário
func BuscarUsuarioCompleto(usuarioID uint64, r *http.Request) (Usuario, error) {

	// criando os canais
	canalUsuario := make(chan Usuario)          // informações do usuário
	canalSeguidores := make(chan []Usuario)     // array de quem seque o usuário
	canalSeguindo := make(chan []Usuario)       // array de quem o usuário segue
	canalPublicacoes := make(chan []Publicacao) // array das publicações do usuário

	// chamando funções usando concorrência
	go buscarDadosDoUsuario(canalUsuario, usuarioID, r)
	go buscarSeguidores(canalSeguidores, usuarioID, r)
	go buscarSeguindo(canalSeguindo, usuarioID, r)
	go buscarPublicacoes(canalPublicacoes, usuarioID, r)

	return Usuario{}, nil
}

func buscarDadosDoUsuario(canal <-chan Usuario, usuarioID uint64, r *http.Request) {

}

func buscarSeguidores(canal <-chan []Usuario, usuarioID uint64, r *http.Request) {

}

func buscarSeguindo(canal <-chan []Usuario, usuarioID uint64, r *http.Request) {

}

func buscarPublicacoes(canal <-chan []Publicacao, usuarioID uint64, r *http.Request) {

}
