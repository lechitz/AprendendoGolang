package main

import (
	"fmt"
	"math"
)

type forma interface { // Interface
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

type retangulo struct { // Struct do Retangulo
	altura float64
	largura float64
}

func (r retangulo) area() float64 { // Método
	return r.altura * r.largura
}

type circulo struct { // Struct do Círculo
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{10,15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
