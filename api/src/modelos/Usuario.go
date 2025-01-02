package modelos

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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

func (usuario *Usuario) validar(etapa string) error {

	tudoCerto := true
	resultado := "Os seguintes campos estão em branco: "

	// verificação de nome
	if usuario.Nome == "" {
		resultado += "nome"
		tudoCerto = false
	}

	// verificação de nick
	if usuario.Nick == "" {

		if !tudoCerto {
			resultado += ", "
		}

		resultado += "nick"
		tudoCerto = false
	}

	// verificação de e-mail
	if usuario.Email == "" {

		if !tudoCerto {
			resultado += ", "
		}

		resultado += "e-mail"
		tudoCerto = false
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	// verificação de senha
	if etapa == "cadastro" && usuario.Senha == "" {

		if !tudoCerto {
			resultado += ", "
		}

		resultado += "senha"
		tudoCerto = false
	}

	// verificando se está tudo correto
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
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.formatar()

	return nil
}
