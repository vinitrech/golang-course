package main

import "fmt"

func main() {

	//---------------------------- ARRAYS INTERNOS ----------------------------
	slice3 := make([]float32, 10, 11) // cria um array com 15 posicoes e retorna um slice com as 10 primeiras

	fmt.Println(slice3, len(slice3), cap(slice3))

	slice3 = append(slice3, 150)

	fmt.Println(slice3, len(slice3), cap(slice3))

	slice3 = append(slice3, 150) // ao estourar as posicoes limites iniciais, dobra o tamanho do array referenciado automaticamente

	fmt.Println(slice3, len(slice3), cap(slice3))

	slice4 := make([]float32, 5) // mantem a capacidade no mesmo valor do tamanho
}
