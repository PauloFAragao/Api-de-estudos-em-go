package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// CarregarTelaDeLogin vai renderizar a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	// renderizando a pagina de login
	utils.ExecutarTemplate(w, "login.html", nil)
}
