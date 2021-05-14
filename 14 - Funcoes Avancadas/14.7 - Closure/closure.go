package main

import "fmt"

func closure() func() {
	texto := "Dentro da funcao closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao

}

func main() {
	texto := "Dentro da funcao main"
	fmt.Println(texto)

	//Embora tenha uma variavel com o mesmo nome da funcao sendo
	//chamada o valor exibido e o original da funcao, como se fosse
	//uma "memoria"
	funcaoNova := closure()
	funcaoNova()
}
