package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá mundo!") // "go" usado para fazer com que o programa continue mesmo sem ter terminado de executar essa linha de comando
	escrever("proxima frase")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}