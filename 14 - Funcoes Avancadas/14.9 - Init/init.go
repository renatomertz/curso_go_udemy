package main

import "fmt"

var n int

//Nao importa a ordem, sempre sera executada em primeiro lugar (aqui antes da main)
//Pode ter um por arquivo
//Usar para setup, iniciar variaveis
func init() {
	fmt.Println("Executando a funcao init")
	n = 10
}

func main() {
	fmt.Println("Executando a funcao main")
	fmt.Println(n)
}
