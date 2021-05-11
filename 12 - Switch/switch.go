package main

import "fmt"

//Testa a mesma variavel no switch
func diaDaSemana(numero int8) string {
	var diaDaSemana string
	switch numero {
	//Nao precisa do break
	case 1:
		diaDaSemana = "Domingo"
		//Faz o codigo cair na proxima condicao (sem avaliar)
		//Pouco utilizado
		fallthrough
	case 2:
		diaDaSemana = "Segunda-feira"
	case 3:
		diaDaSemana = "Terça-feira"
	case 4:
		diaDaSemana = "Quarta-feira"
	case 5:
		diaDaSemana = "Quinta-feira"
	case 6:
		diaDaSemana = "Sexta-feira"
	case 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número inválido"
	}

	return diaDaSemana
}

//Testa variavel no case (pode testar variaveis diferentes)
func diaDaSemana2(numero int8) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-feira"
	case numero == 3:
		return "Terça-feira"
	case numero == 4:
		return "Quarta-feira"
	case numero == 5:
		return "Quinta-feira"
	case numero == 6:
		return "Sexta-feira"
	case numero == 7:
		return "Sábado"
	default:
		return "Número inválido"
	}

}

func main() {
	dia := diaDaSemana(3)
	fmt.Println(dia)

	dia2 := diaDaSemana2(5)
	fmt.Println(dia2)

	fmt.Println("-------------")
	dia3 := diaDaSemana(13)
	fmt.Println(dia3)

	fmt.Println("------fallthrough-------")
	dia4 := diaDaSemana(1)
	fmt.Println(dia4)

}
