package main

import "fmt"

func soma(numeros ...int) (total int) {

	for _, numero := range numeros {
		total += numero
	}

	return
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println(soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	escrever("Vinicius comeu esse tanto de reais: ", 1, 3, 4, 5, 6)
}
