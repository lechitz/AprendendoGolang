package main

import "fmt"

func main()  {
	canal := make(chan string,2) //limite de 2 do canal RX TX

	canal <- "Olá mundo!"
	canal <- "Oie!" //máximo que o canal permite transmitir

	mensagem1 := <- canal
	mensagem2 := <- canal //máximo que o canal permite receber

	fmt.Println(mensagem1)
	fmt.Println(mensagem2)
}