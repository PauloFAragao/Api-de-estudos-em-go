package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// CarregarTelaDeLogin renderiza a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {

	// lendo o cookie
	cookie, _ := cookies.Ler(r)

	// verificando a propriedade token
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	// renderizando a pagina de login
	utils.ExecutarTemplate(w, "login.html", nil)
}

// CarregarPaginaDeCadastroDeUsuario carrega a pagina de cadastro do usuário
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	// renderizando a pagina de login
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

// CarregarPaginaPrincipal carrega a página principal com as publicações
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf("%s/publicacoes", config.APIURL)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	// verificando o status code do html
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var publicacoes []modelos.Publicacao

	// pegando as informações da publicação
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// pegando o cookie para depois pegar o id do usuário
	cookie, _ := cookies.Ler(r)

	// pegando o id
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	// renderizando a pagina de login
	utils.ExecutarTemplate(w, "home.html", struct {
		Publicacoes []modelos.Publicacao
		UsuarioID   uint64
	}{
		Publicacoes: publicacoes,
		UsuarioID:   usuarioID,
	})
}

// CarregarPaginaDeAtualizacaoDePublicacao carrega a página de edição de publicação
func CarregarPaginaDeAtualizacaoDePublicacao(w http.ResponseWriter, r *http.Request) {

	// capturando os parâmetros
	parametros := mux.Vars(r)

	// capturando o id que veio como parâmetro
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// criando a url
	url := fmt.Sprintf("%s/publicacoes/%d", config.APIURL, publicacaoID)

	// requisição
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	// verificando o status code de resposta
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var publicacao modelos.Publicacao

	// tratando o json
	if erro = json.NewDecoder(response.Body).Decode(&publicacao); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "atualizar-publicacao.html", publicacao)

}

// CarregarPaginaDeUsuarios carrega a página com os usuários que atender o filtro passado
func CarregarPaginaDeUsuarios(w http.ResponseWriter, r *http.Request) {

	// capturando os parâmetros
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	// uri da api
	url := fmt.Sprintf("%s/usuarios?usuario=%s", config.APIURL, nomeOuNick)

	//requisição
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	// verificando o status code de resposta
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var usuarios []modelos.Usuario

	// abrindo o json
	if erro = json.NewDecoder(response.Body).Decode(&usuarios); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// mandando renderizar o html
	utils.ExecutarTemplate(w, "usuarios.html", usuarios)
}

// CarregarPerfilDoUsuario carrega a página do perfil do usuário
func CarregarPerfilDoUsuario(w http.ResponseWriter, r *http.Request) {
	// capturando os parâmetros
	parametros := mux.Vars(r)

	// capturando o id que veio como parâmetro
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// pegando o cookie
	cookie, _ := cookies.Ler(r)

	// pegando o id do usuário no cookie
	usuarioLogadoID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	// verificando se o perfil que está sendo acessado é o do usuário logado
	if usuarioID == usuarioLogadoID {
		http.Redirect(w, r, "/perfil", http.StatusFound)
	}

	// pedindo pra api os dados do usuário
	usuario, erro := modelos.BuscarUsuarioCompleto(usuarioID, r)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// renderizando a pagina
	utils.ExecutarTemplate(w, "usuario.html", struct {
		Usuario         modelos.Usuario
		UsuarioLogadoID uint64
	}{
		Usuario:         usuario,
		UsuarioLogadoID: usuarioLogadoID,
	})

}

// CarregarPerfilDoUsuarioLogado carrega a página do perfil do usuário logado
func CarregarPerfilDoUsuarioLogado(w http.ResponseWriter, r *http.Request) {

	// pegando o cookie
	cookie, _ := cookies.Ler(r)

	// pegando o id do usuário no cookie
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	// pedindo pra api os dados do usuário
	usuario, erro := modelos.BuscarUsuarioCompleto(usuarioID, r)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "perfil.html", usuario)

}

// CarregarPaginaDeEdicaoDeUsuario carrega a página para edição dos dados do usuário
func CarregarPaginaDeEdicaoDeUsuario(w http.ResponseWriter, r *http.Request) {

	// pegando o cookie
	cookie, _ := cookies.Ler(r)

	// pegando o id do usuário no cookie
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	// canal para enviar para a função que pesquisa dados do usuário
	canal := make(chan modelos.Usuario)

	// chamando a função que busca dados do usuário usando concorrência
	go modelos.BuscarDadosDoUsuario(canal, usuarioID, r)

	// recebendo os dados do canal
	usuario := <-canal

	// verificando se recebeu os dados corretamente
	if usuario.ID == 0 {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: "erro ao buscar o usuário"})
		return
	}

	utils.ExecutarTemplate(w, "editar-usuario.html", usuario)
}

// CarregarPaginaDeAtualizacaoDeSenha carrega a página para atualização da senha do usuário
func CarregarPaginaDeAtualizacaoDeSenha(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "atualizar-senha.html", nil)
}
