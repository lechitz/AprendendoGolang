package main

import (
	"bytes"
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
	dog := cachorro{"Meimei", "Viralata", 5}
	fmt.Println(dog)

	cachorroEmJSON, erro := json.Marshal(dog)	//convertendo para JSON com o Marchal.
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)

	fmt.Println(bytes.NewBuffer(cachorroEmJSON)) //cria um buffer baseado num slice de bytes

	dog2 := map[string]string {
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorro2EmJSON, erro := json.Marshal(dog2)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorro2EmJSON)
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))

}
