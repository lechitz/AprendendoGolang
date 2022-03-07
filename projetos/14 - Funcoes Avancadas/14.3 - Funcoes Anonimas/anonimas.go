package main

import "fmt"

func main() {

	// IMPRIMIR NA TELA

	func(texto string) { //Função anônima
		fmt.Println(texto)
	}("Aqui entra o parâmetro da função") // colocar () após a chave da função, significa que já quero que execute a função

	// RETORNO

	funcaoRetorno := func(retornar string) string {
		return fmt.Sprintf("Retornou -> %s %d", retornar, 99) //  string %s | numero %d |
	}("oie")

	fmt.Println(funcaoRetorno)

}
