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
// teste6 : 2*100.0 + 100.0 = 300
// teste7 : 100.0 + 100.0 *2 = 400
func Calcula(expressao string) (float64, error) {
	simbolosRegex := regexp.MustCompile("\\+|\\*|\\^|-|%")

	found := simbolosRegex.FindAllIndex([]byte(expressao), -1)

	lastFound := found[0][0]
	op := string(expressao[lastFound])
	digito1X, _ := strconv.ParseFloat(expressao[0:lastFound], 64)
	digito2X, _ := strconv.ParseFloat(expressao[lastFound+1:], 64)
	return retornaCalculo(digito1X, digito2X, op)

	//var operacoes []byte
	//valorTemp := 0
	// for x, y := range found {
	// 	lastFound := y[0]
	// 	digito1X, _ := strconv.ParseFloat(expressao[0:lastFound], 64)
	// 	digito2X, _ := strconv.ParseFloat(expressao[lastFound+1:found[x+1][0]], 64)
	// 	valorTemp, _ := retornaCalculo(digito1X, digito2X, string(expressao[lastFound]))
	// }

	// arrayDividido2 := expressaoRegular.Split(array, -1)
	// fmt.Println(arrayDividido2)
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
