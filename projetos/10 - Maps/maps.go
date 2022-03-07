package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{ // map -> dentro [tipo da chave] / fora []tipo dos valores
		"nome": "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario) //retorna: map[nome:Pedro sobrenome:Silva]
	fmt.Println(usuario["nome"]) //retorna: Pedro

	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro": "Felipe",
			"segundo": "Bravo",
		},
		"curso": {
			"nome": "Engenharia",
			"Universidade": "UFJF",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome") //deleta a chave "nome" de usuario2
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string { //cria uma nova chave e valor para o usuario2
		"signo": "virgem",
	}

	fmt.Println(usuario2)
}