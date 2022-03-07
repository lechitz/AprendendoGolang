package main

import "fmt"

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados) //Usando apenas uma vez o go worker, o processo é lento
	go worker(tarefas, resultados) //Usando apenas uma vez o go worker, o processo é um pouco mais rápido no inicio
	go worker(tarefas, resultados) //Usando apenas uma vez o go worker, o processo é mais rápido ainda..

	for i := 0; i < 45; i++ { //preenchendo o canal de tarefas
		tarefas <- i //enviando os valores
	}
	close(tarefas) //fechar o canal para não receber mais dados

	for i := 0; i < 45; i++ {
		resultado := <- resultados
		fmt.Println(resultado)
	}
}

func worker(tarefas <-chan int, resultados chan<- int) { // <-chan: tarefa só recebe | chan<- : resultados só envia
	for numero := range tarefas { //para cada numero dentro do canal de tarefas...
		resultados <- fibonacci(numero) //calcula o numero de fibonacci e envia para resultados
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}