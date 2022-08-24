package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade int    `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dalmata", 10}

	cachorroJson, err := json.Marshal(c)

	if err != nil {
		log.Fatal("Pobrema")
	}

	fmt.Println(cachorroJson) // slice de bytes
	fmt.Println(bytes.NewBuffer(cachorroJson))

	c2 := map[string]string{
		"nome":  "Toby",
		"raca":  "Poodle",
		"idade": "5",
	}

	c2Json, err := json.Marshal(c2)

	if err != nil {
		log.Fatal("Pobrema")
	}

	fmt.Println(c2Json) // slice de bytes
	fmt.Println(bytes.NewBuffer(c2Json))
}
