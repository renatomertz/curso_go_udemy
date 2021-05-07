package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint16
}

func main() {
	fmt.Println("Arquivos structs")
	var u usuario
	u.nome = "Ranato"
	u.idade = 43
	fmt.Println(u)

	endrecoExemplo := endereco{"Rua dos bobos numero", 0}

	u2 := usuario{"Renato", 43, endrecoExemplo}
	fmt.Println(u2)

	u3 := usuario{nome: "Renato"}
	fmt.Println(u3)

	u4 := usuario{idade: 43}
	fmt.Println(u4)
}
