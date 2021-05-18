package main

import "fmt"

//Pode receber qualquer coisa
//Qualquer coisa mesmo!!!!
//Pode ser parametro, retorno, tipo , chave/valor de map
//Usar com muito cuidado pois isso burla a tipagem forte do Go
func generica(inter interface{}) {
	fmt.Println(inter)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	//Exemplo de um uso sadio desse recurso
	fmt.Println(1, 2, "string", false, true, float64(123456))

	//Quizumba, bagunca!!!!
	mapa := map[interface{}]interface{}{
		"aaa":           "bbb",
		123:             456,
		float32(123465): 1,
	}

	fmt.Println(mapa)
}
