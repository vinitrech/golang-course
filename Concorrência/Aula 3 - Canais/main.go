package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá mundo", canal)

	// for {
	// 	mensagem, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for i := range canal {
		fmt.Println(i)
	}
}

func escrever(param string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- param
		time.Sleep(time.Second)
	}

	close(canal)
}
