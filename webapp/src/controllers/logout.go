package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// FazerLogout remove os dados de autenticação salvos no browser do usuário
func FazerLogout(w http.ResponseWriter, r *http.Request) {
	// mandando o cookie em branco
	cookies.Deletar(w)

	// mandando o usuário para a pagina de login
	http.Redirect(w, r, "/login", http.StatusFound)
}
