package main

import (
	"html/template"
	"net/http"
	"curso-alura/models"
)


var views = template.Must(template.ParseGlob("views/*.html"))


func main(){
	http.HandleFunc("/", index)
	http.ListenAndServe(":5000", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	produtos := models.BuscaTodosOsProdutos()
	views.ExecuteTemplate(w, "Index", produtos)
}

