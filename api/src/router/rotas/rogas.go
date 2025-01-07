package rotas

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rodas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar colocar todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {

	rotas := rotasUsuarios                     // /usuarios
	rotas = append(rotas, rotaLogin)           // /login
	rotas = append(rotas, rotasPublicacoes...) // /publicacoes

	// configurando todas as rotas
	for _, rota := range rotas {

		// verificando se a rota precisa de autenticação
		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)

		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}

	}

	return r
}
