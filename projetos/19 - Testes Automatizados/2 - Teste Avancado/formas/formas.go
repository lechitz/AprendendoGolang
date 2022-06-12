package formas

import (
	"math"
)

type Forma interface { // Interface
	Area() float64
}

type Retangulo struct { // Struct do Retangulo
	Altura float64
	Largura float64
}

func (r Retangulo) Area() float64 { // Método
	return r.Altura * r.Largura
}

type Circulo struct { // Struct do Círculo
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}