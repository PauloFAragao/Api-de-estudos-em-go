package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa um usuário da rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (usuario *Usuario) validar() error {

	tudoCerto := true
	resultado := "Os seguintes campos estão em branco: "

	if usuario.Nome == "" {
		resultado += "nome"
		tudoCerto = false
	}

	if usuario.Nick == "" {

		if !tudoCerto {
			resultado += ", "
		}

		resultado += "nick"
		tudoCerto = false
	}

	if usuario.Email == "" {

		if !tudoCerto {
			resultado += ", "
		}

		resultado += "e-mail"
		tudoCerto = false
	}

	if usuario.Senha == "" {

		if !tudoCerto {
			resultado += ", "
		}

		resultado += "senha"
		tudoCerto = false
	}

	if !tudoCerto {
		return errors.New(resultado)
	} else {
		return nil
	}

}

func (usuario *Usuario) formatar() {

	//retirando os espaços das extremidades das strings
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}

	usuario.formatar()

	return nil
}
