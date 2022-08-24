package main

import "fmt"

func main() {
	numero := 10
	numero2 := 20

	switch numero {
	case 10:
		fmt.Println(numero)
	default:
		fmt.Println("Deu bosta patrão")
	}

	switch {
	case numero2 > 10:
		fmt.Println("deu certo patrao, é maior")
		fallthrough // cai direto na proxima condicao!
	default:
		fmt.Println("Deu ruim patrao, é menor q 10")
	}
}
