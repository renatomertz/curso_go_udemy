package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Heranca")

	p1 := pessoa{"Renato", "Mertz", 43, 1.80}
	fmt.Println(p1)

	e1 := estudante{p1, "Ciencia da computação", "Anhembi Morumbi"}
	//Pode imprimir total ou...
	fmt.Println(e1)
	//...somente um atributo do struct 'PAI'
	//Poderia ser e1.pessoa.nome
	fmt.Print(e1.nome)
}
