package modelos

// DadosAutenticacao contém  o id e o token go usuário autenticados
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
