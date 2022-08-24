package main

import "fmt"

func init() { // sempre e executada antes da main
	fmt.Println("Funcao init")
}

func main() {
	fmt.Println("Funcao main")
}
