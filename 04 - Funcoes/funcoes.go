package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//Quando tem varios parametros do mesmo tipo ele sempre pega o tipo do ultimo
func calculosMatematico(n1, n2 int8) (int8, int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	subtracao2 := n2 - n1
	return soma, subtracao, subtracao2
}

func main() {
	soma := somar(1, 2)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto função 1")
	println(resultado)

	resultadoSoma, resultadoSubtracao, _ := calculosMatematico(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	//Quando tem mais de um retorno e nao se quer usar todos basta colocar o _ (underline) no lugar do que seria o retorno

	_, _, resultadoSubtracao2 := calculosMatematico(10, 15)
	fmt.Println(resultadoSubtracao2)
}
