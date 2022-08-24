package enderecos

import (
	"fmt"
	"strings"
	"testing"
)

type cenarioTeste struct {
	endereco string
	retorno  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Run("Enderecos válidos", func(t *testing.T) {
		testes := []cenarioTeste{
			{"Rua 1", "válido"},
			{"Avenida 1", "válido"},
			{"Estrada 1", "válido"},
			{"Rodovia 1", "válido"},
		}

		for i, teste := range testes {
			fmt.Println("Teste ", i)
			retornoDaFuncao := strings.Split(TipoDeEndereco(teste.endereco), " ")[1]

			if retornoDaFuncao != teste.retorno {
				t.Error("Erro no teste ", i)
			}
		}
	})

	t.Run("Enderecos válidos", func(t *testing.T) {
		testes := []cenarioTeste{
			{"abc", "inválido"},
			{" ", "inválido"},
			{".", "inválido"},
			{"!@#", "inválido"},
		}

		for i, teste := range testes {
			fmt.Println("Teste ", i)
			retornoDaFuncao := strings.Split(TipoDeEndereco(teste.endereco), " ")[1]

			if retornoDaFuncao != teste.retorno {
				t.Error("Erro no teste ", i)
			}
		}
	})
}
