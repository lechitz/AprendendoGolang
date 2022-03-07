package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá Mundo!"), escrever("Programando em go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<- canal)
	}
}

func multiplexar(canalDeEntrada1, canalDeEntrada2 <- chan string) <- chan string { //Criamos dois canais que só recebem valores
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
				case mensagem := <- canalDeEntrada1: //Caso tenha uma mensagem disponível no canalDeEntrada1...
					canalDeSaida <- mensagem //Jogamos a mensagem no canalDeSaida
				case mensagem := <- canalDeEntrada2:
					canalDeSaida <- mensagem
			}
		}
	}()

	return canalDeSaida //dessa forma centralizamos a comunicação de dois canais dentro de um canal só
}

func escrever(texto string) <-chan string { //return do chan só recebe dados
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %v", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000))) //tempo pseudo-aleatório: escolhe 1 numero de 0 a 2000
		}
	}()

	return canal
}