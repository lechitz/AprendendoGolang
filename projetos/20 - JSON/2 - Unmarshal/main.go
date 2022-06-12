package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json:"nome"`
	Raca string `json:"raca"`
	Idade uint 	`json:"idade"`
}

func main() {
	cachorroEmJSON := `{"nome":"Meimei","raca":"Viralata","idade":5}`

	var dog cachorro

	//passamos com & para representar o endereço de memória - para que "dog" fique alterada depois dessa linha
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &dog) ; erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(dog)

	cachorro2EmJSON := `{"nome": "Toby","raca": "Poodle"}`

	dog2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &dog2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(dog2)
}
