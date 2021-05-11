package main

import (
	"fmt"
)

func main() {
	//i := 0
	//Como se fosse o do / while
	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando o i")
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	//Mais comum
	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando o j -", j)
	// 	time.Sleep(time.Second)
	// }

	//Clausula range
	nomes := [3]string{"João", "Davi", "Lucas"}

	//Retorna o indice e o valor armazenado no array, sempre nessa ordem
	//for indice, nome := range nomes {
	//Pode ignorar o primeiro passando o underline
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	//Retorna o codigo ASCII, precisa converter para string
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//Loop infinito
	//for {

	//}
}
