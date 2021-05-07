package main

import "fmt"

func main() {
	//Aritimeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	//Nao e possivel fazer operacoes com tipos diferentes
	//O trecho abaixo da erro
	//var numero1 int16 = 10
	//var numero2 int32 = 10
	//soma2 := numero1 + numero2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// Fim aritmeticos

	//Atribuicao

	//Com igual onde precisa indicar o tipo da variavel
	var variavel string = "String"
	//Por inferencia onde o tipo vem do que foi atribuido
	variavel2 := "String2"

	fmt.Println(variavel, variavel2)

	//Fim atribuicao

	//Relacionais
	//Sempre retornam bool
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	//Fim relacionais

	//Operadores logicos
	fmt.Println("---------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//Fim operadores logicos

	//Operadores unarios
	//No go nao tem pre-fixado (++numero ou --numero)
	numero := 10
	numero++
	numero += 15 //numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20
	fmt.Println(numero)

	numero *= 3
	fmt.Println(numero)

	numero /= 10
	fmt.Println(numero)

	numero %= 3
	fmt.Println(numero)

	//Nao ter ternario no go, premissa de as coisas serem o mais simples possivel
}
