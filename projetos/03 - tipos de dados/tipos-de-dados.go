package main

import (
	"errors"
	"fmt"
)

func main() {
	 var numero int64 = -10000000
	 fmt.Println(numero)

	 var numero2 uint32 = 10000
	 fmt.Println(numero2)

	 //alias
	 // INT32 = RUNE
	 var numero3 rune = 1234
	 fmt.Println(numero3)

	 // BYTE = UINT8
	 var numero4 byte = 123
	 fmt.Println(numero4)

	 var numeroReal1 float32 = 123.45
	 fmt.Println(numeroReal1)

	 var numeroReal2 float64 = 1234454.22
	 fmt.Println(numeroReal2)

	 numeroReal3 := 1547.112
	 fmt.Println(numeroReal3)

	 var erro error = errors.New("Erro interno")
	 fmt.Println(erro)
}