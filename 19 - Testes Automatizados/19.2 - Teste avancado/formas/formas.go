package formas

import (
	"math"
)

//A interface com o metodo a ser implementado
type Forma interface {
	area() float64
	perimetro() float64
	aumentarArea(n uint) float64
}

//Para implementar a interface tem que atender o requisito:
//1 - Criar um metodo para o struct com o mesmo NOME e RETORNO
//que tem a INTERFACE
func (r Retangulo) Area() float64 {
	return r.altura * r.largura
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func (r Retangulo) Perimetro() float64 {
	return 2*r.altura + 2*r.largura
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func (r Retangulo) AumentarArea(n uint) float64 {
	return float64(n) * r.Area()
}

func (c Circulo) AumentarArea(n uint) float64 {
	return float64(n) * c.Area()
}

type Retangulo struct {
	largura float64
	altura  float64
}

type Circulo struct {
	raio float64
}
