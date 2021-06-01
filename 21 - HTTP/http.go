package main

import (
	"log"
	"net/http"
)

func home(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Ola mundo!"))
}

func usuarios(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Carregar página de usuarios"))
}

func main() {

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	//SOBE O SERVIDOR
	log.Fatal(http.ListenAndServe(":5000", nil))
}
