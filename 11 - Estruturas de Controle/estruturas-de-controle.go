package main

import "fmt"

func main() {
	numero := -15
	//No go tem que ter as chaves e nem na mesma linha
	//Pode ser if numero > 15 {	fmt.Println("Maior que 15")	} Mas nao e usual
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor igual a 15")
	}

	//If init
	//variavel outroNumero fica limitada ao escopo do if, depois nao esta mais acessivel
	if outroNumero := numero; outroNumero > 0 {
		fmt.Printf("Numero e maior do que zero")
	} else if outroNumero < -10 {
		fmt.Println("Numero e menor que zero")
	} else {
		fmt.Println("Entre 0 e -10")
	}
}
