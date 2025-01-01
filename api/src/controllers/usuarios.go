package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
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

	//validando os dados inseridos pelo usuário
	if erro = usuario.Preparar(); erro != nil {
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
	w.Write([]byte("Buscando um usuário!"))
}

// AtualizarUsuario atualiza os dados de um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário!"))
}

// DeletarUsuario deleta um usuário do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}
