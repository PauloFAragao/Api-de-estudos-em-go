package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"webapp/src/respostas"
)

// FazerLogin utiliza o e-mail e senha do usuário para autenticar na aplicação
func FazerLogin(w http.ResponseWriter, r *http.Request) {
	// pegando o corpo da requisição
	r.ParseForm()

	// montando json para enviar para api com um map
	usuario, erro := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// enviando o json para a api
	response, erro := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	// temporário
	token, _ := ioutil.ReadAll(response.Body)

	fmt.Println(response.StatusCode, string(token))
}
