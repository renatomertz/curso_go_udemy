package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

//o metodo esta atrelado a alguma estrutura, nesse caso o struct
//mas pode ser com uma inrtaface
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuario %s no banco de dados.\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuario1", 17}
	usuario1.salvar()

	usuario2 := usuario{"Davi", 30}
	usuario2.salvar()

	maiorDeIdade := usuario1.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario1.fazerAniversario()
	maiorDeIdade = usuario1.maiorDeIdade()
	fmt.Println(maiorDeIdade)

}
