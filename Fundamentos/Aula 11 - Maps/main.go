package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Vinicius",
		"sobrenome": "Tonello",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"chave": {
			"nome":      "Vinicius",
			"sobrenome": "Tonello",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "chave") // remove chave do map

	usuario2["signo"] = map[string]string{ //insere uma nova chave ao map
		"nome": "Gemeos",
	}
}
