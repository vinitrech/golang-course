package main

import "fmt"

func inverterSinal(numero *int) {
	*numero *= -1
}

func main() {
	numero1 := 40

	fmt.Println(numero1)

	inverterSinal(&numero1)

	fmt.Println(numero1)
}
