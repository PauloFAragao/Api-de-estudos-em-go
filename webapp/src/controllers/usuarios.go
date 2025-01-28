package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
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

// PararDeSeguirUsuario chama a api para parar de seguir um usuário
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {

	// pegando os parâmetros
	parametros := mux.Vars(r)

	//pegando o usuário Id do parâmetro
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// url para da api
	url := fmt.Sprintf("%s/usuarios/%d/parar-de-seguir", config.APIURL, usuarioID)

	// fazendo requisição na api
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// tratando o status code da api
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

// SeguirUsuario chama a api para seguir um usuário
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	// pegando os parâmetros
	parametros := mux.Vars(r)

	//pegando o usuário Id do parâmetro
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// url para da api
	url := fmt.Sprintf("%s/usuarios/%d/seguir", config.APIURL, usuarioID)

	// fazendo requisição na api
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// tratando o status code da api
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

// EditarUsuario chama a api para editar um usuário
func EditarUsuario(w http.ResponseWriter, r *http.Request) {

	// pegando o corpo da requisição
	r.ParseForm()

	// montando json para enviar para api com um map
	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// pegando o cookie
	cookie, _ := cookies.Ler(r)

	// pegando o id do usuário do cookie
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	//url para a api
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioID)

	// recebendo resposta da api
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(usuario))
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

// AtualizarSenha chama a api para atualizar a senha do usuário
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	// pegando o corpo da requisição
	r.ParseForm()

	// montando json para enviar para api com um map
	senhas, erro := json.Marshal(map[string]string{
		"atual": r.FormValue("atual"),
		"nova":  r.FormValue("nova"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// pegando o cookie
	cookie, _ := cookies.Ler(r)

	// pegando o id do usuário do cookie
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	//url para a api
	url := fmt.Sprintf("%s/usuarios/%d/atualizar-senha", config.APIURL, usuarioID)

	// recebendo resposta da api
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(senhas))
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

// DeletarUsuario chama a api para deletar o usuário
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

	// pegando o cookie
	cookie, _ := cookies.Ler(r)

	// pegando o id do usuário do cookie
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	//url para a api
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioID)

	// recebendo resposta da api
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodDelete, url, nil)
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
