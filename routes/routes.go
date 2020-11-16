package routes

import (
	"net/http"
	"curso-alura/controllers"
)

func CarregaRotas(){
	http.HandleFunc("/", controllers.Index)
}
