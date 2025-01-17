package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
)

// CriarPublicacao chama a API para cadastrar uma publicação no banco de dados
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	// pegando o corpo da requisição
	r.ParseForm()

	publicacoes, erro := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// url para api
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)

	// enviando o json para a api
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacoes))
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

// CurtirPublicacao chama a API para curtir uma publicação
func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	// capturando os parametros
	parametros := mux.Vars(r)

	// id da publicação
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// criando a rota para enviar para a api
	url := fmt.Sprintf("%s/publicacoes/%d/curtir", config.APIURL, publicacaoID)

	// fazendo a requisição para a api
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	// verificando o status code da resposta
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}

// DescurtirPublicacao chama a API para descurtir uma publicação
func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	// capturando os parametros
	parametros := mux.Vars(r)

	// id da publicação
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// criando a rota para enviar para a api
	url := fmt.Sprintf("%s/publicacoes/%d/descurtir", config.APIURL, publicacaoID)

	// fazendo a requisição para a api
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	// verificando o status code da resposta
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}
