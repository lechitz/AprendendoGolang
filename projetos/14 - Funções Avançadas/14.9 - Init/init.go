package main

import "fmt"

var n int //criando como variável global

func init() { // Faz com que seja executada antes da função main, independente da ordem de onde estiver a init na cascata
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Executando a função main")
	fmt.Println(n)
}