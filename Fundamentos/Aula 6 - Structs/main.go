package main

import "fmt"

type usuario struct {
	nome            string
	idade           uint8
	enderecoExemplo endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var u usuario
	u.nome = "Vinicius"
	u.idade = 23

	fmt.Println(u.nome, u.idade)

	u2 := usuario{"Vinicius", 23, endereco{"logradouro", 25}}

	fmt.Println(u2)

	u3 := usuario{nome: "Vinicius"}

	fmt.Println(u3)
}
