package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Publicacoes representa um repositório de publicações
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um repositório de publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar Insere uma publicação no banco de dados
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	// Query
	statement, erro := repositorio.db.Prepare(
		"Insert Into publicacoes  (titulo, conteudo, autor_id) values (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	//executando a query
	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	// capturando o id da publicação
	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

// BuscarPorID trás uma única publicação do bando de dados
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {

	// Query
	linha, erro := repositorio.db.Query(`
		Select p.*, u.nick 
		From publicacoes p 
		Inner Join usuarios u On u.id = p.autor_id
		Where p.id = ?`, publicacaoID)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao modelos.Publicacao

	// executando a query
	if linha.Next() {
		// capturando os dados e jogando em publicacao
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

// Buscar trás as publicações dos usuários seguidos e também do próprio usuário que fez a requisição
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	// query -- Distinct -- para não vir resultados duplicados
	linhas, erro := repositorio.db.Query(`
		Select Distinct p.*, u.nick 
		From publicacoes p 
		Inner Join usuarios u on u.id = p.autor_id 
		Inner Join seguidores s on p.autor_id = s.usuario_id
		Where p.id = ? Or s.seguidor_id = ?
		Order by 1 desc `, usuarioID, usuarioID) // Order by p.criadaEm -- eu acho que ordenar por criadaEm faz mais sentido
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()
	// Order by p.criadaEm

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		// capturando os dados e jogando em publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Atualizar altera os dados de uma publicação no banco de dados
func (repositorio Publicacoes) Atualizar(pulicacaoID uint64, publicacao modelos.Publicacao) error {
	// query
	statement, erro := repositorio.db.Prepare("Update publicacoes Set titulo = ?, conteudo = ? Where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	// executando
	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, pulicacaoID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui uma publicação do banco de dados
func (repositorio Publicacoes) Deletar(pulicacaoID uint64) error {
	// query
	statement, erro := repositorio.db.Prepare("Delete From publicacoes Where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	// executando
	if _, erro = statement.Exec(pulicacaoID); erro != nil {
		return erro
	}

	return nil
}
