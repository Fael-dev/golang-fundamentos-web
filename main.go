package main

import (
	"database/sql"
	"html/template"
	"net/http"
	_ "github.com/lib/pq"
)

type Produto struct {
	Nome string
	Descricao string
	Preco float64
	Quantidade int
}

var views = template.Must(template.ParseGlob("views/*.html"))


func main(){
	db := conexaoBd()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":5000", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	produtos := []Produto{
		{ Nome:"Camiseta", Descricao:"Azul", Preco:39, Quantidade: 3 },
		{"Fone", "Preto", 25, 10 },
		{"Headset", "Preto", 50, 10 },
		{"Novo Produto", "Red", 20, 15 },
	}

	views.ExecuteTemplate(w, "Index", produtos)
}
func conexaoBd() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=postgres host=localhost sslmode=disabled"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
