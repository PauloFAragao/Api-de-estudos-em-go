package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Representa um repositório de usuários
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (u Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	return 0, nil
}
