package main

import "fmt"

func recuperarExecucao() {
	fmt.Println("Tentativa de recuperar a execução")
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada")
	}
}

func alunoAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A MÉDIA É EXATAMENTE 6") // o panic ao ser executado, chama todas as funcoes defer
}

func main() {
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("Pós Execução")
}
