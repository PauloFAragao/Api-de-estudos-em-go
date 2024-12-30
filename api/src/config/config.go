package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexaoBanco é a string de conexão com o MySql.
	StringConexaoBanco = ""

	// Porta onded a API vai estar rodando
	Porta = 0
)

// Carregar vai inicializar as variáveis de ambiente
func Carregar() {
	var erro error

	// godotenv vai ler o arquivo .env que contem as configurações
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	// converte string em int
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000 // porta default caso não consiga ler o arquivo
	}

	// string de conexão
	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

}
