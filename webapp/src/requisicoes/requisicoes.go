package requisicoes

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

// FazerRequisicaoComAutenticacao é utilizada para colocar o token na requisição
func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {

	//criando a requisição
	request, erro := http.NewRequest(metodo, url, dados)
	if erro != nil {
		return nil, erro
	}

	// lendo o cookie
	cookie, _ := cookies.Ler(r)

	// adicionando o token ao request
	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	// criando o client
	client := &http.Client{}

	// fazendo a requisição
	response, erro := client.Do(request)
	if erro != nil {
		return nil, erro
	}

	return response, nil

}
