package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// CriarUsuario cria um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	// lendo o corpo do http
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario

	// convertendo o json em usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	// Conexão com o banco dedados
	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}
	//defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	repositorio.Criar(usuario)

}

// BuscarUsuarios busca todos os usuários do banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários!"))
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
