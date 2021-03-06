package controllers

import (
	"curso-alura/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var views = template.Must(template.ParseGlob("views/*.html"))

func Index(w http.ResponseWriter, r *http.Request){
	produtos := models.BuscaTodosOsProdutos()
	views.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request){
	views.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversao do preco: ", err)
		}

		qtdConvertido, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversao da quantidaed: ", err)
		}

		models.CriarProduto(nome, descricao, precoConvertido, qtdConvertido )
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request){
	idProduto := r.URL.Query().Get("id")
	models.DeletarProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request){
	idProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idProduto)
	views.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertido, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Erro na conversao do ID para Int.")
		}

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversao do Preco para Float.")
		}

		quantidadeConvertido, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversao da Quantidade para Int.")
		}

		models.AtualizaProduto(idConvertido, nome, descricao, precoConvertido, quantidadeConvertido)
	}

	http.Redirect(w, r, "/", 301)


}