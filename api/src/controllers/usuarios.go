package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CriarUsuario cria um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	// lendo o corpo do http
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {

		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario

	// convertendo o json em usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {

		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// validando os dados inseridos pelo usuário
	if erro = usuario.Preparar("cadastro"); erro != nil {

		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// Conexão com o banco dedados
	db, erro := banco.Conectar()
	if erro != nil {

		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {

		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios busca todos os usuários do banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	// pegando a string pesquisada
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	// conexão com o banco dedados
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	// criando um repositório de usuários
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	// mandando pesquisar no banco
	usuarios, erro := repositorio.Buscar(nomeOuNick)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)

}

// BuscarUsuario busca um usuário por id no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	// capturando os parâmetros
	parametros := mux.Vars(r)

	// convertendo o parâmetro para inteiro
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// conexão com o banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	// passando o banco de dados para o repetitório
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	// buscando usuário
	usuario, erro := repositorio.BuscarPorID(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuario)

}

// AtualizarUsuario atualiza os dados de um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	// capturando os parâmetros
	parametros := mux.Vars(r)

	// capturando o Id
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// pegando o id que veio no token
	usuarioIdNoToken, erro := autenticacao.ExtrairUsuarioId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	// verificando se o usuario tem permissão para fazer a alteração
	if usuarioID != usuarioIdNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar um usuário que não seja o seu"))
		return
	}

	// capturando os dados do corpo da requisição
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {

		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	// struct para os dados do usuario
	var usuario modelos.Usuario

	// extraindo dados do json
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// verificando se os dados recebidos são validos
	if erro = usuario.Preparar("edicao"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// conexão com o banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	// criando um repositório
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Atualizar(usuarioID, usuario); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

// DeletarUsuario exclui as informações de um usuário do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	// capturando os parâmetros
	parametros := mux.Vars(r)

	// capturando o Id
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// pegando o id que veio no token
	usuarioIdNoToken, erro := autenticacao.ExtrairUsuarioId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	// verificando se o usuario tem permissão para fazer a alteração
	if usuarioID != usuarioIdNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível deletar um usuário que não seja o seu"))
		return
	}

	// conexão com o banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	// criando um repositório
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Deletar(usuarioID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// SeguirUsuario permite que um usuario siga outro
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	// pegando o id do seguidor
	seguidorId, erro := autenticacao.ExtrairUsuarioId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	// capturando os parâmetros
	parametros := mux.Vars(r)

	// convertendo em int
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// verificando se o usuario quer seguir a si mesmo
	if seguidorId == usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível seguir a você mesmo"))
		return
	}

	// conexão com o banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	// repositório
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	// mandando seguir
	if erro = repositorio.Seguir(usuarioID, seguidorId); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// PararDeSeguirUsuario permite que um usuario pare de seguir outro
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	// pegando o id do seguidor
	seguidorId, erro := autenticacao.ExtrairUsuarioId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	// capturando os parâmetros
	parametros := mux.Vars(r)

	// convertendo em int
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// verificando se o usuario quer deixar de seguir a si mesmo
	if seguidorId == usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível parar seguir a você mesmo"))
		return
	}

	// conexão com o banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	// repositório
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	// mandando parar de seguir
	if erro = repositorio.PararDeSeguir(usuarioID, seguidorId); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// BuscarSeguidores trás todos os seguidores de um usuario
func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {

	// capturando os parâmetros
	parametros := mux.Vars(r)

	// convertendo em int
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// conexão com o banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	// repositório
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	// capturando os seguidores
	seguidores, erro := repositorio.BuscarSeguidores(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, seguidores)
}

// BuscarSeguindo trás todos os usuarios que um determinado usuário está seguindo
func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {

	// capturando os parâmetros
	parametros := mux.Vars(r)

	// convertendo em int
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// conexão com o banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	// repositório
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	// capturando quem segue
	usuarios, erro := repositorio.BuscarSeguindo(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

// AtualizarSenha permite alterar a senha de um usuario
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {

	// pegando o id que veio no token
	usuarioIdNoToken, erro := autenticacao.ExtrairUsuarioId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	// capturando os parâmetros
	parametros := mux.Vars(r)

	// convertendo em int
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// verificando se o usuario no token e o usuario nos parametros são diferentes
	if usuarioIdNoToken != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar a senha de um usuário que não seja o seu"))
		return
	}

	// lendo o corpo da requisisão
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	// struct para trocar a senha
	var senha modelos.Senha

	// retirando as senhas do json
	if erro = json.Unmarshal(corpoRequisicao, &senha); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// conexão com o banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	// repositório
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	//buscando a senha salva no banco
	senhaSalvaNoBanco, erro := repositorio.BuscarSenha(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	// verificando se a senha fornecida como senha atual e a senha gardada no banco são iguais
	if erro = seguranca.VerificarSenha(senhaSalvaNoBanco, senha.Atual); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, errors.New("a senha atual não condiz com a senha que está salva no banco"))
		return
	}

	// inserindo hash na nova senha
	senhaComHash, erro := seguranca.Hash(senha.Nova)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// inserindo a senha no banco
	if erro = repositorio.AtualizarSenha(usuarioID, string(senhaComHash)); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}
