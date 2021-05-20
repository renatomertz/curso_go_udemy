package main

import "fmt"

//Comoo se fosse uma fila de processos
//e varios workers pegando dessa fila
func main() {
	const tamanho int = 40

	tarefas := make(chan int, tamanho)
	resultados := make(chan int, tamanho)

	//Nao tem problema chamar antes de popular
	//porque ao chegar na funcao vai ficar
	//esperando o canal receber um valor
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < tamanho; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := 1; i <= tamanho; i++ {
		resultado := <-resultados
		fmt.Println(i, "-", resultado)
	}

}

// <-chan = Canal que so recebe
// chan<- = Canal que so envia
//Dessa forma deixam de ser bidirecionais
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
