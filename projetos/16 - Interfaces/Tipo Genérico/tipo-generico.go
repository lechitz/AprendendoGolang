package main

import "fmt"

func generica(interf interface{}) { //dessa maneira aceitaria qualquer coisa
	fmt.Println(interf)
}

func main() {
	generica("string")
	generica(1)
	generica(true)

	mapa := map[interface{}] interface{} {
		1: "string",
		"string": "string",
		true: float64(12),
	}

	fmt.Println(mapa)
}
