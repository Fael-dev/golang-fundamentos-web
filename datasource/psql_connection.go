package datasource

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConexaoBd() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}