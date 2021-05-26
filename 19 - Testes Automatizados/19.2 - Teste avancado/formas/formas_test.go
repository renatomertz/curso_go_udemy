package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	//Uado para separar o que está sendo testado
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{12, 10}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		//Fatalf interrompe a execucao dos testes
		if areaEsperada != areaRecebida {
			t.Fatalf("Valor recebido %f diferente de valor esperado %f", areaEsperada, areaRecebida)
		}

		perimetroEsperado := float64(44)
		perimetroRecebido := ret.Perimetro()

		if perimetroEsperado != perimetroRecebido {
			t.Fatalf("Valor recebido %f diferente de valor esperado %f", perimetroEsperado, perimetroRecebido)
		}

		areaAumetadaEsperada := areaEsperada * 10
		areaAumentadaRecebida := ret.AumentarArea(10)

		if areaAumetadaEsperada != areaAumentadaRecebida {
			t.Fatalf("Valor recebido %f diferente de valor esperado %f", areaAumetadaEsperada, areaAumentadaRecebida)
		}

	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := math.Pi * 100
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Valor recebido %f diferente de valor esperado %f", areaEsperada, areaRecebida)
		}

		perimetroEsperado := math.Pi * 20
		perimetroRecebido := circ.Perimetro()

		if perimetroEsperado != perimetroRecebido {
			t.Fatalf("Valor recebido %f diferente de valor esperado %f", perimetroEsperado, perimetroRecebido)
		}

		areaAumetadaEsperada := areaEsperada * 10
		areaAumentadaRecebida := circ.AumentarArea(10)

		if areaAumetadaEsperada != areaAumentadaRecebida {
			t.Fatalf("Valor recebido %f diferente de valor esperado %f", areaAumetadaEsperada, areaAumentadaRecebida)
		}

	})

}
