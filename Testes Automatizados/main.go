package main

import (
	"fmt"
	"testes/enderecos"
	/*Imports com . antes indicam que é o pacote principal, não precisando chamar pacote.funcaotal()*/)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("avenida dos retardados mentais")

	fmt.Println(tipoEndereco)
}

/*
go test ./... > testa todos os arquivos de teste de todos os pacotes
go test --cover > indica a porcentagem de cobertura da função dos testes
go test --coverprofile teste.txt > salva o relatorio de teste contendo quais partes não foram testadas
go tool cover --func=cobertura.txt > aponta quais funções não foram cobertas a partir do relatorio
go tool cover --html=cobertura.txt > mostra um arquivo html apontando o que não foi coberto
*/
