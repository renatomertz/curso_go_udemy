package servidor

import (
	"crud/banco"
	"fmt"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

//Funcao que cria um usuario
func CriarUsuario(rw http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		rw.Write([]byte("Falha ao ler o corpo da requisicao"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(body, &usuario); erro != nil {
		rw.Write([]byte("Erro ao converter usuario para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		rw.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO usuarios (NOME, EMAIL) VALUES (?, ?)")
	if erro != nil {
		rw.Write([]byte("Erro ao preparar o statement"))
		return
	}
	defer statement.Close()

	insert, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		rw.Write([]byte("Erro ao executar o statement"))
		return
	}

	id, erro := insert.LastInsertId()
	if erro != nil {
		rw.Write([]byte("Erro ao obter id do usuario"))
		return
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(fmt.Sprintf("Usuario isnsrido com sucesso! ID: %d", id)))

}
