package controllers

import (
	"curso-alura/models"
	"html/template"
	"net/http"
)

var views = template.Must(template.ParseGlob("views/*.html"))

func Index(w http.ResponseWriter, r *http.Request){
	produtos := models.BuscaTodosOsProdutos()
	views.ExecuteTemplate(w, "Index", produtos)
}
