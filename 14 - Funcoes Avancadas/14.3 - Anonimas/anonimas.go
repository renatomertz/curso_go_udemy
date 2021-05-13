package main

import "fmt"

func main() {

	retorno := func(texto string) string {
		return fmt.Sprintf("Recibo -> %s", texto)
	}("Ola mundo")

	fmt.Printf(retorno)
}
