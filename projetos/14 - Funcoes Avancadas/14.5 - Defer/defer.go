package main

import "fmt"

func funcao1() {
	fmt.Println("Função 1")
}

func funcao2() {
	fmt.Println("Função 2")
}

func funcao3() {
	fmt.Println("Função 3")

}

func funcao4() {
	fmt.Println("Função 4")
}

func main() {
	funcao1()
	defer funcao2() // o DEFER faz com que a chamada seja levada diretamente para o final das iterações (adiada para final)
	funcao3()
	funcao4()
}
