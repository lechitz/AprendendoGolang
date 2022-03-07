package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil { //se não tivermos um recover nesse programa, o programa morre e nada depois disso vai ser executado
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MEDIA É EXATAMENTE 6!") //vai chamar todas as funções que tem defer.
}

func main() {
	fmt.Println(alunoEstaAprovado(9, 6))
	fmt.Println("Pós Execução")
}
