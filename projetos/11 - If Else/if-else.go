package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 16

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 { //if init -> cria a variável dentro do if, e não pode ser usada depois fora
		fmt.Println("Número é maior que zero!")
	} else {
		fmt.Println("Número é menor ou igual a zero")
	}
}