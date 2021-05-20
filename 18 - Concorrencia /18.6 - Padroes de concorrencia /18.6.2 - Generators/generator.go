package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Ola Mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

//Encapsula a chamada de uma goroutine e devolve um canal
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
