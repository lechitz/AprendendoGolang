package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao

}

func main() {
	soma := somar(10,20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt) //imprime uma vez
		return txt
	}

	resultado := f("Texto da função")
	fmt.Println(resultado) //imprime pela segunda vez

	resultadoSoma, _ := calculosMatematicos(10,15) //o underline substitui a outra variavel da função, para caso eu não queira usar ela aqui.
	fmt.Println(resultadoSoma)
}