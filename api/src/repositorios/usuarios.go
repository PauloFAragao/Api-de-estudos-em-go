package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

// Representa um repositório de usuários
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuário no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	// preparando string com a query
	statement, erro := repositorio.db.Prepare(
		"Insert Into usuarios (nome, nick, email, senha) Values (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	// Executando a string com a query no banco de dados
	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	// pegando o id onde o usuário foi inserido
	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar traz todos os usuários que atendem um filtro de nome ou nick
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOuNick%

	// query
	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome LIKE ? OR nick LIKE ? ",
		nomeOuNick, nomeOuNick,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	// slice de usuarios
	var usuarios []modelos.Usuario

	// adicioando os usuarios ao slice
	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarPorID traz um usuário do banco de dados
func (repositorio Usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	// Query
	linhas, erro := repositorio.db.Query(
		"Select id, nome, nick, email, criadoEm From usuarios Where id = ?", ID,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, nil
		}
	}

	return usuario, nil
}

// Atualizar altera as informações de um usuário no banco de dados
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	// query para atualizar
	statemente, erro := repositorio.db.Prepare(
		"Update usuarios Set nome = ?, nick = ?, email = ? Where id = ? ")
	if erro != nil {
		return erro
	}
	defer statemente.Close()

	// executando a query
	if _, erro = statemente.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui as informações de um usuário no banco de dados
func (repositorio Usuarios) Deletar(ID uint64) error {

	// query para deletar
	statemente, erro := repositorio.db.Prepare("Delete From usuarios Where id = ?")
	if erro != nil {
		return erro
	}
	defer statemente.Close()

	// executando a query
	if _, erro = statemente.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorEmail busca um usuário por email e retorna o seu id e senha com hash
func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	// Criando a query
	linha, erro := repositorio.db.Query("Select id, senha From usuarios Where email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	// executar a query
	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil

}

// Seguir permite que um usuario siga outro
func (repositorio Usuarios) Seguir(usuarioID, seguidorId uint64) error {

	// query
	statement, erro := repositorio.db.Prepare(
		"Insert Ignore Into seguidores (usuario_id, seguidor_id) Values (? ,?)", // o ignore é pra ignorar caso o dado já esteja no banco de dados
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorId); erro != nil {
		return erro
	}

	return nil
}

// PararDeSeguir permite que um usuario pare de seguir o outro
func (repositorio Usuarios) PararDeSeguir(usuarioID, seguidorId uint64) error {

	// query
	statement, erro := repositorio.db.Prepare(
		"Delete From seguidores Where usuario_id = ? And seguidor_id = ? ",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorId); erro != nil {
		return erro
	}

	return nil

}

// BuscarSeguidores trás todos os seguidores de um usuario
func (repositorio Usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	// query
	linhas, erro := repositorio.db.Query(`
		Select u.id, u.nome, u.nick, u.email, u.criadoEm
		From usuarios u 
		Inner Join seguidores s On u.id = s.seguidor_id
		Where s.usuario_id = ? 
		`, usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	//  slice de usuarios que vai receber a resposta do banco
	var usuarios []modelos.Usuario

	// executando a query
	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarSeguindo trás todos os usuários que um determinado usuário está seguindo
func (repositorio Usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {

	//query
	linhas, erro := repositorio.db.Query(`
		Select u.id, u.nome, u.nick, u.email, u.criadoEm
		From usuarios u
		Inner Join seguidores s On u.id = s.usuario_id
		Where s.seguidor_id = ?`, usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	//  slice de usuarios que vai receber a resposta do banco
	var usuarios []modelos.Usuario

	// executando a query
	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}

// BuscarSenha traz a senha de um usuário pelo ID
func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	//query
	linha, erro := repositorio.db.Query("Select senha From usuarios Where id = ?", usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	// executando a query
	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

// AtualizarSenha altera a senha de um usuário
func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	//query
	statement, erro := repositorio.db.Prepare("Update usuarios Set senha = ? Where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	// executando a query
	if _, erro = statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}

	return nil
}
