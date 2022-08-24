package main

import "fmt"

func main() {
	canal := make(chan string, 2) // caso estoure o limite de envios de dados, causa deadlock

	canal <- "Olá mundo"
	canal <- "Olá mundo 2"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem, mensagem2)
}
