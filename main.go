package main

import (
	"html/template"
	"net/http"
)

var views = template.Must(template.ParseGlob("views/*.html"))

func main(){
	http.HandleFunc("/", index)
	http.ListenAndServe(":5000", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	views.ExecuteTemplate(w, "Index", nil)
}