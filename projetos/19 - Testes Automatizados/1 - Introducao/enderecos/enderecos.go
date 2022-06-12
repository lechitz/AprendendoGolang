package enderecos

import "strings"

// TipoDeEndereco verifica se um endereço tem um tipo valido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "travessa", "estrada", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0] //o índice [0] retorna apenas a primeira posição
	// Exemplo de entrada: "Rua Itamar Franco"
	// Exemplo de saída: ["Rua", "Itamar", "Franco"]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco) //strings.Title deixa a primeira letra maiúscula de cada palavra
	}

	return "Tipo Invalido"
}