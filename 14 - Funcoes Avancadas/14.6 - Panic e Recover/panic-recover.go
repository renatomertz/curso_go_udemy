package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Recuperado com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	media := (n1 + n2) / 2
	defer recuperarExecucao()

	if media == 6 {
		//Se entrar nessa condição o programa para a execução
		//entrando em pânico
		//Antes de parar chama as funcoes defer
		panic("A MEDIA E EXATAMENTE 6")
	}
	return media > 6
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Executado")
}
