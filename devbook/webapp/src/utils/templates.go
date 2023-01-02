package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// CarregarTemplates Insere templates html na variavel templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

//ExecutarTemplate renderiza a pagona na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
