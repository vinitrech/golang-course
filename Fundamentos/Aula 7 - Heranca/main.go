package main

import "fmt"

type pessoa struct {
	nome string
}

type estudante struct {
	pessoa
	curso string
}

func main() {

	p1 := pessoa{"Vinicius"}
	p2 := estudante{pessoa{"Vinicius"}, "Curso"}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p2.nome)

}
