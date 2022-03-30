package main

import (
	"fmt"
	"introducao-testes/Enderecos"
)

func main() {
	tipoEndereco := Enderecos.TipoDeEndereco("Rodovia Paulista")
	fmt.Println(tipoEndereco)
}