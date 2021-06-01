package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func home(rw http.ResponseWriter, r *http.Request) {
	u := usuario{
		"Joao",
		"joao.pedro@gmail.com",
	}
	templates.ExecuteTemplate(rw, "home.html", u)
}

func main() {
	templates = template.Must(template.ParseGlob(("*.html")))
	http.HandleFunc("/home", home)

	//SOBE O SERVIDOR
	fmt.Println("Executando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
