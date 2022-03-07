package main //os métodos estão associados a uma estrutura, ao invés de algo solto

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

func (u usuario) salvar() { //Um método é uma função que possui uma struct associada!
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool { //cria uma função "maiorDeIdade" que é associada a struct "usuário" através da variável "u" que criamos
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() { //faço a referência por ponteiro para mudar o valor na alocação de memória da variável
	u.idade++
}

func main() {
	usuario1 := usuario{"Felipe", 30}
	usuario1.salvar()

	usuario2 := usuario{"Silvana", 58}
	usuario2.salvar()

	maiordeIdade1 := usuario1.maiorDeIdade()
	maiorDeIdade2 := usuario2.maiorDeIdade()
	fmt.Println(maiordeIdade1)
	fmt.Println(maiorDeIdade2)
	usuario1.fazerAniversario() //chama a função de aniversário e incrementa 1
	fmt.Println(usuario1.idade)
	fmt.Println(usuario2.idade)
}
