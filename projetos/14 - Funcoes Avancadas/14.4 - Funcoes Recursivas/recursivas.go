package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("\nFunções Recursivas")
	posicao := uint(15)

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}
	fmt.Print("A soma dos dois ultimos é: ")
	fmt.Println(fibonacci(posicao))
	fmt.Printf("Ao todo são: %d posições\n", posicao)
}
