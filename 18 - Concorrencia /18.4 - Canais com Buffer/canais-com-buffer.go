package main

import "fmt"

func main() {

	//Se criar desse jeito vai dar deadlock
	//Deadlock esta enviando o valor mas nao tem quem recebe
	//canal := make(chan string) -- Aqui trava a execucao esperando alguem receber a resposta

	//O canal com buffer so vai bloquear o trafego de dados
	//quando o canal atingir a capacidade maxima
	canal := make(chan string, 2) //Canal com buffer de tamanho 2
	canal <- "Olá mundo"
	canal <- "Eu disse ola"
	//canal <- "Respondam" -- Esse daria deadlock
	mensagem := <-canal
	mensagem2 := <-canal
	//mensagem3 := <-canal -- deadlock porque nao tem canal enviando
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
