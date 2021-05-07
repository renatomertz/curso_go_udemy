package main

import (
	"errors"
	"fmt"
)

func main() {
	//INICIO numeros INTEIROS

	//4 tipos de ints, int8, int16, int32 e int64
	//Quanto maior o numero maior o tamanho dele
	//Se declarar sem o numero (int) ele vai assumir a arquitetura do computador,
	//se 64 bits sera equivalente a int64, se 32 bits int32
	//Também suportam numero negativos
	var numero int8 = 127
	fmt.Println(numero)

	//Somente numeros positivos, sem sinal
	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//alias, equivale a int32
	var numero3 rune = 12345
	fmt.Println(numero3)

	//alias, equivale a uint8
	var numero4 byte = 255
	fmt.Println(numero4)
	//TERMINO numeros INTEIROS

	//INICIO numeros REAIS

	//Numeros reais
	//2 tipos, float32 e float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12389798789789.45
	fmt.Println(numeroReal2)

	//Nao pode declarar somente float (sem o tamanho)
	//Na inferencia ele vai seguir a arquitetura, 32 ou 64 bits
	numeroReal3 := 12340.98
	fmt.Println(numeroReal3)

	//TERMINO numeros REAIS

	//INICIO STRING
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//Npo GO não tem char, todas string tem que ser colocada entre aspas duplas
	//Se colocar entre aspas simples será impresso impresso o número que representa o caracter na tabela ASCII
	char := 'B'
	fmt.Println(char)

	//FIM STRING

	//INICIO VALOR 0
	//Todo tipo tem seu valor inicial, o valor 0 caso seja declarada uma variavel sem ser atribuido valor
	//Para inteiro e float e 0 e para string e vazio
	//Com declaracao por inferencia nao funciona porque ten que inicializar com valor
	var intZero int16
	fmt.Println(intZero)

	var strZero string
	fmt.Println(strZero)

	//Outros tipos
	//Boolean (valor 0 e false)
	var booleano1 bool = true
	fmt.Println(booleano1)

	//Erro
	//Valor inicial <nill> (null)
	//Para inicializar precisa usar usar um pacote nativo do go (errors)
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
