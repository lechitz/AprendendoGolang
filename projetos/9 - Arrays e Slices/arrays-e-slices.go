package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1[5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição1", "Posição2", "Posição3", "Posição4", "Posição5"}
	fmt.Println(array2)

	array3 := [...]int{1, 26, 14, 58, 207} // usar [...] para quando não quiser limitar o numero do array. Ele cria um array do tamanho informado aqui e nada além.
	fmt.Println(array3)


	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18) //vai adicionar uma posicao a mais no slice ao final.
	fmt.Println(slice)

	slice2 := array2[1:3] //esse slice vai pegar uma fatia do array2, indo da posicao 1 até a 3
	fmt.Println(slice2)

	array2[1] = "PosiçãoAlterada"
	fmt.Println("----------")
	fmt.Println(array2)
	fmt.Println(slice2)

	fmt.Println("----------")
	slice3 := make([]float32, 10, 15) //tamanho de 10 e capacidade máxima de 15
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) //capacidade
}