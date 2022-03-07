package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<- canal) //Imprime o valor que está chegando no canal
	}
}

func escrever(texto string) <-chan string { //return do chan só recebe dados
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %v", texto)
			time.Sleep(time.Millisecond * 500) //time de 0,5 segundos
		}
	}()

	return canal
}