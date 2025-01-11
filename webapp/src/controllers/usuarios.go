package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/respostas"
)

// CriarUsuario chama a api para cadastrar um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	// pegando o corpo da requisição
	r.ParseForm()

	// montando json para enviar para api com um map
	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// url para api
	url := fmt.Sprintf("%s/usuarios", config.APIURL)

	// enviando o json para a api
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	// verificando se o status code está no range de erro
	if response.StatusCode >= 400 {
		// enviando a resposta de erro
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	// enviando resposta
	respostas.JSON(w, response.StatusCode, nil)
}
