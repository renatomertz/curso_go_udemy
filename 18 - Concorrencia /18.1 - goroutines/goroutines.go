package main

import (
	"fmt"
	"time"
)

func main() {
	// A funcao escrever vai rodar infinitamente na primeira chamada
	// mas a go routine (go) diz para o programa executar e nao esperar
	// o termino da execuxao para seguir o programa
	// Se colocra gouroutime em todas as chamadas nao dara tempo de
	// executar nada (para testar colocar go na frente das duas chamadas)
	go escrever("Ola mundo!")
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
