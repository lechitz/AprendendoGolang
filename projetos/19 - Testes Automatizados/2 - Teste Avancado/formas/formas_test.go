package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T)  {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10,12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Area recebida %f é diferente da Area esperada %f", areaRecebida,areaEsperada)
		}
	})
	
	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(100*math.Pi)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Area recebida %f é diferente da Area esperada %f",areaRecebida,areaEsperada)
		}
	})
	
}