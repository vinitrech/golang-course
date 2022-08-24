package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá Mundo")
	escrever("Apaga isso sao nudes")
}

func escrever(param string) {
	for {
		fmt.Println(param)
		time.Sleep(time.Second)
	}
}
