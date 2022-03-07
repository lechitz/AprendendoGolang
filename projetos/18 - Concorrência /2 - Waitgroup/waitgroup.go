package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup //pacote padrão do Go

	waitGroup.Add(2) // Quantidade de goroutines rodando ao mesmo tempo - grupo de espera

	go func() { //função anônima, chamo ela com a ''go'' - goroutines
		escrever("texto 1")
		waitGroup.Done() // Esse done conta 1 vez o delta do .Add
	}()

	go func() {
		escrever("texto 2")
	}()

	waitGroup.Wait() //espera todas as goroutines serem contadas antes de executar
}

func escrever(texto string) {
	for i:=0; i<5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}