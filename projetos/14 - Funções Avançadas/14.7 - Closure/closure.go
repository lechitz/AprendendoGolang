package main

import "fmt"

func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Print(texto)
	}

	return funcao
}

func main() {
	texto := "dentro do main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}