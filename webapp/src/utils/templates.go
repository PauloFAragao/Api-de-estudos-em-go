package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// CarregarTemplates insere os templates html na variável templates
func CarregarTemplates() {
	// carregando os arquivos html na variável
	templates = template.Must(templates.ParseGlob("views/*.html"))

	// carregando os templates
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// ExecutarTemplate renderiza uma página html na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
