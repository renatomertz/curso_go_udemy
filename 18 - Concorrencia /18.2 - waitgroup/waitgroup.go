package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Maneira de sicronizar as go routines (a menos utilizada)
	var waitGroup sync.WaitGroup
	//Especifica a quntidade de go routines que vão
	// executar ao mesmo tempo
	waitGroup.Add(4)

	go func() {
		escrever("Ola mundo!")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Go routine 3!")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Go routine 4!")
		waitGroup.Done() // -1
	}()

	//Indica que o programa espere para finalizar apos as
	//duas funcoes terminarem
	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
