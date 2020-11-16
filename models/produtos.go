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
