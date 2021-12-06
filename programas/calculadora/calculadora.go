package calculadora

import (
	"errors"
	"math"
	"regexp"
	"strconv"
)

// teste1 : "2+2"
// teste2" "10.5-5.5"
// teste3 : "1000.0*3"
// teste4 : 2^8
// teste5 : 100%5
// teste6 : 2*100.0 + 100.0 = 300  => nao podem existir espacos na expressao
// teste7 : 100.0 + 100.0 *2 = 400
func Calcula(expressao string) (float64, error) {
	simbolosRegex := regexp.MustCompile("\\+|\\*|\\^|-|%|\\/")

	matrizOperacoes := simbolosRegex.FindAllIndex([]byte(expressao), -1)

	if len(matrizOperacoes) == 0 {
		return 0.0, errors.New("Colocar expressao pf")
	}
	indicePrimeiraOperacao := matrizOperacoes[0][0]
	numeroEsquerda, _ := strconv.ParseFloat(expressao[0:indicePrimeiraOperacao], 64)
	var err error
	numeroDireita := 0.0
	for indice, posicaoOperacao := range matrizOperacoes {
		indiceOperacaoAtual := posicaoOperacao[0]
		possuiProximaOperacao := len(matrizOperacoes) > indice+1
		if possuiProximaOperacao {
			indiceProximaOperacao := matrizOperacoes[indice+1][0]
			numeroDireita, _ = strconv.ParseFloat(expressao[indiceOperacaoAtual+1:indiceProximaOperacao], 64)
		} else {
			numeroDireita, _ = strconv.ParseFloat(expressao[indiceOperacaoAtual+1:], 64)
		}
		numeroEsquerda, err = realizaOperacaoMatematica(numeroEsquerda, numeroDireita, string(expressao[indiceOperacaoAtual]))
	}

	return numeroEsquerda, err
}

func realizaOperacaoMatematica(digito1, digito2 float64, operacao string) (float64, error) {

	switch operacao {
	case "+":
		{
			return digito1 + digito2, nil
		}
	case "-":
		{
			return digito1 - digito2, nil
		}
	case "/":
		{
			return digito1 / digito2, nil
		}
	case "%":
		{
			return float64(int(digito1) % int(digito2)), nil
		}
	case "*":
		{
			return digito1 * digito2, nil
		}
	case "^":
		{
			return math.Pow(digito1, digito2), nil
		}
	default:
		{
			return 0.0, errors.New("Coloca a operacao ai cara!")

		}
	}
}
