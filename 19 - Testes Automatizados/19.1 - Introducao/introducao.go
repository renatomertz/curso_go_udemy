package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	enderecoResidencial := "Rua dos bobos num 0"
	fmt.Println(enderecos.TipoDeEndereco(enderecoResidencial))
}
