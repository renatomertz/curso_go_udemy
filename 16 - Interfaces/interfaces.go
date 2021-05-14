package main

import (
	"fmt"
	"math"
)

//A interface com o metodo a ser implementado
type forma interface {
	area() float64
	perimetro() float64
	aumentarArea(n uint) float64
}

//Funcao que recebe uma forma e mostra a area
func escreverArea(f forma) {
	fmt.Printf("A area da forma e %0.2f\n", f.area())
}

//Para implementar a interface tem que atender  requisito:
//1 - Criar um metodo para o struct com o mesmo NOME e RETORNO
//que tem a INTERFACE
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func escreverPerimetro(f forma) {
	fmt.Printf("O perimetro da forma e %0.2f\n", f.perimetro())
}

func (r retangulo) perimetro() float64 {
	return 2*r.altura + 2*r.largura
}

func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func escreverOTamanhoDaAreaAumentado(f forma, n uint) {
	fmt.Printf("O tamanho da area aumentado em %d ficou em %0.2f\n", n, f.aumentarArea(n))
}

func (r retangulo) aumentarArea(n uint) float64 {
	return float64(n) * r.area()
}

func (c circulo) aumentarArea(n uint) float64 {
	return float64(n) * c.area()
}

type retangulo struct {
	largura float64
	altura  float64
}

type circulo struct {
	raio float64
}

func main() {
	ret := retangulo{10, 15}
	escreverArea(ret)
	escreverPerimetro(ret)
	escreverOTamanhoDaAreaAumentado(ret, 10)

	circ := circulo{10}
	escreverArea(circ)
	escreverPerimetro(circ)
	escreverOTamanhoDaAreaAumentado(circ, 10)
}
