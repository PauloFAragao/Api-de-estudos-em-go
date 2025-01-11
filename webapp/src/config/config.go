package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL apresenta a URL para comunicação com a API
	APIURL = ""

	// Porta onde a aplicação web está rodando
	Porta = 0

	// HashKey é utilizada para autenticar o cookie
	HashKey []byte

	// BlockKey é utilizada para criptografar os dados do cookie
	BlockKey []byte
)

// Carregar inicializa as variáveis de ambiente
func Carregar() {
	var erro error

	// lendo arquivo .env
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	// lendo a porta
	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	// lendo as outras configurações
	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))

}
