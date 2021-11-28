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
	simbolosRegex := regexp.MustCompile("\\+|\\*|\\^|-|%")

	found := simbolosRegex.FindAllIndex([]byte(expressao), -1)

	// lastFound := found[0][0]
	// fmt.Println(lastFound)
	// op := string(expressao[lastFound])
	// digito1X, _ := strconv.ParseFloat(expressao[0:lastFound], 64)
	// digito2X, _ := strconv.ParseFloat(expressao[lastFound+1:], 64)
	// return retornaCalculo(digito1X, digito2X, op)

	if len(found) == 0 {
		return 0.0, errors.New("Colocar expressao pf")
	}

	valorTemp, _ := strconv.ParseFloat(expressao[0:found[0][0]], 64)
	var er error
	digito2 := 0.0
	for x, y := range found {
		nextFound := y[0]
		if len(found) > x+1 {
			digito2, _ = strconv.ParseFloat(expressao[nextFound+1:found[x+1][0]], 64)
		} else {
			digito2, _ = strconv.ParseFloat(expressao[nextFound+1:], 64)
		}
		valorTemp, er = retornaCalculo(valorTemp, digito2, string(expressao[nextFound]))
	}

	return valorTemp, er
}

func retornaCalculo(digito1, digito2 float64, operacao string) (float64, error) {

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
