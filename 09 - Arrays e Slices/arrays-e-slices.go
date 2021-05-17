package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println()

	var array1 [5]string
	array1[0] = "Posicao1"
	fmt.Println(array1)

	array2 := [5]string{"Posicao1", "Posicao2", "Posicao3", "Posicao4", "Posicao5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//E dinamico pode aumentar ou diminuir
	//Slices sao pedacos, fatias de arrays
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	//Indice 1 e inclusivo (pega a posicao 1)
	//Indice 3 e exclusivo (nao pega a posicao 3)
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posicao alterada"
	fmt.Println(slice2)

	//Arrays Internos
	fmt.Println(("-------------"))
	//Cria um array do tipo float32 com tamanho 6 (ultimo parametro), cria o slice
	//apontando para as 10 primeiras posicoes (primeiro parametro)
	slice3 := make([]float32, 5, 6)
	fmt.Println(("Apos criacao"))
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	//Quando o tamanho do array interno e estourado o go cria um novo com o dobro da
	//capacidade + 1 e aponta o slice
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(("Apos estouro"))
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	//Pode criar sem passar a capacidade
	//Vai assumir o mesmo do tamanho
	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
