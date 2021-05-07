package auxiliar

import "fmt"

//Para poder ser exportada (ou importada) em outro pacotes a função SEMPRE tem que ser escrita com a primeira letra maiuscula
//Letra minuscula indica que ela só pode ser usada internamente
//Precisa de um comentário quando é 'publico'
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	//Está no mesmo pacote então tudo bem.
	escrever2()
}
