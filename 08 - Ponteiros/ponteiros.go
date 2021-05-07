package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")
	//Copia do valor
	variavel1 := 10
	variavel2 := variavel1

	fmt.Println(variavel1, variavel2)

	variavel2++
	fmt.Println(variavel1, variavel2)

	//Ponteiro e uma referencia de memoria
	var variavel3 int
	var ponteiro *int
	fmt.Println(variavel3, ponteiro)

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) //desreferenciacao

	variavel3 = 150

	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) //desreferenciacao
}
