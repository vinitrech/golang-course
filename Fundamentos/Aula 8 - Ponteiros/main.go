package main

import "fmt"

func main() {
	var variavel1 int = 10
	var ponteiro *int = &variavel1

	fmt.Println(variavel1, *ponteiro, ponteiro)

	variavel1 = 50

	fmt.Println(*ponteiro)

}
