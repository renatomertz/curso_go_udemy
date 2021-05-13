package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	posicao := uint(13)
	fmt.Println(fibonacci(posicao))
	fmt.Println(12 * 12)
}
