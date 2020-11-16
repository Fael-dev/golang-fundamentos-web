package main

import (
	"net/http"
	"curso-alura/routes"
)


func main(){
	routes.CarregaRotas()
	http.ListenAndServe(":5000", nil)
}


