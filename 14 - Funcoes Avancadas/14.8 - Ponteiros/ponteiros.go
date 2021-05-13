package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	//Por copia
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numero)
	fmt.Println(numeroInvertido)

	//Por referencia
	novoNumero := numero * 2
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
