package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá mundo"), escrever("Programando em go!"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

//Junta dois ou mais canais em um so
func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)
	go func() {
		for {
			select {
			case mensagem := <-canalEntrada1:
				canalSaida <- mensagem
			case mensagem := <-canalEntrada2:
				canalSaida <- mensagem
			}
		}
	}()
	return canalSaida
}

//Encapsula a chamada de uma goroutine e devolve um canal
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
