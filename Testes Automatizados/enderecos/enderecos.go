package enderecos

import (
	"strings"
)

//TipoDeEndereco retorna se o endereco é válido ou não
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	primeiraPalavra := strings.Split(strings.ToLower(endereco), " ")[0]

	valido := false

	for _, posicao := range tiposValidos {
		if posicao == primeiraPalavra {
			valido = true
		}
	}

	if !valido {
		return "Endereço inválido - " + strings.Title(primeiraPalavra)
	} else {
		return "Endereço válido - " + strings.Title(primeiraPalavra)
	}
}
