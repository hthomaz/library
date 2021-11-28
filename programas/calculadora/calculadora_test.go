package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcula_Soma(t *testing.T) {
	entrada := "2+2"
	saidaEsperada := 4.0
	saida, _ := Calcula(entrada)

	assert.Equal(t, saidaEsperada, saida)
}
func TestCalcula_Subtracao(t *testing.T) {
	entrada := "10.5-5.5"
	saidaEsperada := 5.0
	saida, _ := Calcula(entrada)

	assert.Equal(t, saidaEsperada, saida)
}
func TestCalcula_Multiplicacao(t *testing.T) {
	entrada := "1000.0*3"
	saidaEsperada := 3000.0
	saida, _ := Calcula(entrada)

	assert.Equal(t, saidaEsperada, saida)
}
func TestCalcula_Potencia(t *testing.T) {
	entrada := "2^8"
	saidaEsperada := 256.0
	saida, _ := Calcula(entrada)

	assert.Equal(t, saidaEsperada, saida)
}
func TestCalcula_Modulo(t *testing.T) {
	entrada := "100%5"
	saidaEsperada := 0.0
	saida, _ := Calcula(entrada)

	assert.Equal(t, saidaEsperada, saida)
}

// func TestCalcula_MultiplasOperacoes(t *testing.T) {
// 	entrada := "2*100.0 + 100.0"
// 	saidaEsperada := 0.0
// 	saida, _ := Calcula(entrada)

// 	assert.Equal(t, saidaEsperada, saida)
// }
