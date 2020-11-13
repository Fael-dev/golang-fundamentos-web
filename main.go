package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome string
	Descricao string
	Preco float64
	Quantidade int
}

var views = template.Must(template.ParseGlob("views/*.html"))

func main(){
	http.HandleFunc("/", index)
	http.ListenAndServe(":5000", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	produtos := []Produto{
		{ Nome:"Camiseta", Descricao:"Azul", Preco:39, Quantidade: 3 },
		{"Fone", "Preto", 25, 10 },
		{"Headset", "Preto", 50, 10 },
	}

	views.ExecuteTemplate(w, "Index", produtos)
}