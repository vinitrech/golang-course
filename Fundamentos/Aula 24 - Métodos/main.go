package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("Salvando user", u.nome, u.idade)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Vinicius", 23}
	usuario1.salvar()
	fmt.Println(usuario1.maiorDeIdade(), usuario1.idade)
	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}
