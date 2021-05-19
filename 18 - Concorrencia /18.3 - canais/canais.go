package main

import (
	"fmt"
	"time"
)

//Canal tem duas operacoes: ENVIAR e RECEBER
//Um canal pode estar aberto ou fechado (propriedade de um canal)
//Aberto -  ainda pode receber ou enviar dados
//Fechado - nao pode receber ou enviar dados
//Essas operacoes bloqueiam a execucao do programa
func main() {
	canal := make(chan string)
	//1 - Ao chamar a funcao e passar o canal
	//por conta da clausula go o programa nao vai esperar o termino
	//da execucao da funcao
	go escrever("Ola mundo!", canal)

	fmt.Println("Apos chamada da funcao escrever")

	//Deadlock (all goroutines are asleep - deadlock)
	//Como esta num for infinito ao terminar de executar a
	//funcao escrever o canal continuara esperando receber um valor
	//CUIDADO - nao e pego em tempo de compilacao, so em execucao
	for {
		//Esperendo que o canal receba um valor
		// 3 - Aqui esta obrigando a espera o canal receber valor
		// (enviado pela funcao escrever)
		//mensagem := <-canal Nao espera o retorno do estado do canal
		//e pode dar deadlock
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	canal2 := make(chan string)
	go escrever("Meu canal!", canal2)

	//Forma mais clean de escrever o for do canal
	//Para cada mensagem recebida no canal enquanto
	//ele estiver aberto, faca
	for mensagem := range canal2 {
		fmt.Println(mensagem)
	}

	fmt.Println("Termino do programa")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		//Mandando um valor para dentro do canal
		//2 - Retorna, escreve o valor
		canal <- texto
		time.Sleep(time.Second)
	}
	//Fecha o canal
	close(canal)
}
