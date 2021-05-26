package enderecos

import (
	"strings"
)

// TipoDeEndereco verifica se um endereco tem um tipo valido e o restorna
func TipoDeEndereco(endereco string) string {
	enderecoMinusculo := enderecoEmLetraMinuscula(endereco)
	primeiraPalavraEndereco := primeiraPalavraDoEndereco(enderecoMinusculo)
	return tipoEncontrado(primeiraPalavraEndereco)
}

func enderecoEmLetraMinuscula(endereco string) string {
	return strings.ToLower(endereco)
}

func primeiraPalavraDoEndereco(endereco string) string {
	return strings.Split(endereco, " ")[0]
}

func tipoEncontrado(primeiraPalavra string) string {
	tipoEncontrado := "Invalido"
	for _, tipo := range retornaTiposValidos() {
		if saoPalavrasIguais(primeiraPalavra, tipo) {
			tipoEncontrado = tipo
		}
	}
	return strings.Title(tipoEncontrado)
}

func saoPalavrasIguais(primeiraPalavra, tipo string) bool {
	return primeiraPalavra == tipo
}

func retornaTiposValidos() []string {
	return []string{"rua", "avenida", "estrada", "rodovia"}
}

//Funcao para mostrar a falta de cobertura
func dummy(num1, num2 int8) string {
	if num1 > num2 {
		return "Num1 maior"
	} else if num1 < num2 {
		return "Num1 menor"
	} else {
		return "Num1 igual"
	}
}
