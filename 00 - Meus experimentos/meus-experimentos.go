package main

import "fmt"

func main() {
	var p *[10]int = new([10]int)  // aloca a estrutura da fatia; * p == nulo; raramente útil
	var v []int = make([]int, 100) // a fatia v agora se refere a uma nova matriz de 100 ints

	fmt.Println("Com new", *p)
	fmt.Println("Com make", v)

	(*p)[5] = 10
	fmt.Println("Apos atribuir ao ponteiro", *p)

	v[99] = 20
	fmt.Println("Apos atribuir ao array", v)
}
