package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	for j := 0; j < 10; j += 1 {
		fmt.Println("Incrementando j -> ", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"Felipe", "João", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}
}
