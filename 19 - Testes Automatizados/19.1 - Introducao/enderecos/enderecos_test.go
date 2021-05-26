package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Estrada do Jaboatao", "Estrada"},
		{"Praca da Moca", "Invalido"},
		{"RUA ABC", "Rua"},
		{"AVENIDA ABC", "Avenida"},
		{"", "Invalido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if cenario.retornoEsperado != tipoDeEnderecoRecebido {
			t.Errorf("Tipo invalido. Esperado %s recebido %s", cenario.retornoEsperado, tipoDeEnderecoRecebido)
		}
	}
}

func TestDummy(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Error("Falhou")
	}
}

//A parte de testes e toda nativa do go
//Para rodar de todos os pacotes na raiz go test ./...
//Para ser mais verboso pode usar o go test -v
//Para rodar os testes em paralelo colocar na primeira linha
////t.Parallel() (colocar em duas ou mais que se deseja o paralelismo)
//Para testar a cobertura go test --cover
//Para gerar relatorio de cobertura go test --coverprofile nomeDoArquivo.txt
//Para interpretar melhor o arquivo gerado go tool cover --func=nomeDoArquivo.txt -- Das as funcoes e %
//Para interpretar melhor o arquivo gerado go tool cover --html=nomeDoArquivo.txt -- Abre html mais visual

//Nome do pacote pode conter _test apos o nome do pacote (enderecos_test)
//Mas com isso precisa importar o pacote que está sendo testado
//pode usar alias na declaracao do paco para nao precisar referenciar
////o pacote toda vez (. "introducao-testes/enderecos")
//Se nome do pacote for igual nao precisa importar
//O importe com . (ponto) e mais usado em testes
