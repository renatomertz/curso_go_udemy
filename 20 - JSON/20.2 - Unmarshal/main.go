package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

//Para omitir a tag pode colocar um - (menos)
//Vale para o marshal e unmarshal
type dummy struct {
	Nome  string `json:"-"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJSON := `{"nome":"Rex","raca":"Dalmata","idade":3}`

	var c cachorro
	//A funcao precisa receber um slice de bytes
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	cachorro2EmJSON := `{"nome":"Toby","raca":"Poodle"}`
	c2 := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)

	var d dummy
	//A funcao precisa receber um slice de bytes
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &d); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(d)

}
