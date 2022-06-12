package enderecos_test

import (
	"testing"
	. "introducao-testes/enderecos"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste {
		{"Rua ABC", "Rua"},
		{"Avenida Itamar", "Avenida"},
		{"Rodovia BR040", "Rodovia"},
		{"Travessa Albergue", "Travessa"},
		{"RUA DOS BOBOS", "Rua"}, //Case Sensitive
		{"", "Tipo Invalido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido '%s' e diferente do esperado '%s'.",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}

