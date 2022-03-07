package main

import "fmt"

type usuario struct {
	nome string
	idade int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario
	u.nome = "Felipe"
	u.idade = 29
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua A", 1}

	usuario2 := usuario{"Pedro", 22, enderecoExemplo}
	fmt.Println(usuario2)
}

