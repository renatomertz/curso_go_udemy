package main

import "fmt"

func main() {
	//Forma explicita
	var variavel1 string = "variavel 1"

	//Forma implicita, sabe-se atraves do valor atribuido
	//Inferencia de tipo
	variavel2 := "variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//Declarar varias variaveis
	var (
		variavel3 string = "lalala"
		variavel4 string = "lalala"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel 5", "variavel 6"

	fmt.Println(variavel5, variavel6)

	//Constante - seguem os mesmos modos de criacao da variavel com a diferenca que nao pode alterar o valor
	const constante1 string = "constante 1"
	fmt.Println(constante1)

	//Inverter as variaveis sem o auxilio de uma auxiliar
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}
