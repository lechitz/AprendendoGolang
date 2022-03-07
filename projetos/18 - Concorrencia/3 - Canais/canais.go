package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá mundo!", canal) //depende do canal

	fmt.Println("Depois da função escrever começar")

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	//// MESMO PRINCÍPIO DO FOR ACIMA
	//for {
	//	mensagem, aberto := <- canal //passa para as variáveis o valor que tem em canal
	//	if !aberto {
	//		break //se o canal não tiver aberto, sai do Loop
	//	}
	//	fmt.Println(mensagem)
	//}

	fmt.Println("Fim")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}