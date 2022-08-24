package main

import "fmt"

func main() {
	a := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Ola Cabesao de batatinha")

	fmt.Println(a)
}
