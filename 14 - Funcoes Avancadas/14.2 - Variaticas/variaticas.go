package main

import "fmt"

//Pode receber 0 a 'n' valores
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

//Pode estar junto com outros parametros mas...
//...Tem que ser o ultimo e...
//... nao pode ter dois variaticos
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 6, 200, 102, 12, 13)
	fmt.Println(totalDaSoma)

	escrever("Ola mundo", 1, 2, 3)
}
