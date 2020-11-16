package models

import (
	"curso-alura/datasource"
)

type Produto struct {
	Id int
	Nome string
	Descricao string
	Preco float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := datasource.ConexaoBd()
	selectProdutos, err := db.Query("select * from produtos ")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriarProduto(nome, descricao string, preco float64, quantidade int){
	db := datasource.ConexaoBd()

	inseriDados, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	inseriDados.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletarProduto(id string){
	db := datasource.ConexaoBd()

	deletarProduto, err := db.Prepare("delete from produtos where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(id)

	defer db.Close()
}

func EditaProduto(id string) Produto{
	db := datasource.ConexaoBd()

	produtoBD, err := db.Query("select * from produtos where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoBD.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoBD.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}

	defer db.Close()

	return produtoParaAtualizar
}
