package ordena

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenaSlice_Com1UnicoValor(t *testing.T) {
	entrada := [][]int{{2, 50}}
	saidaEsperada := [][]int{{2, 50}}
	saida := OrdenaSlice(entrada)

	assert.Equal(t, saidaEsperada, saida)
}

func TestOrdenSlice_Com2ValoresConcatenados(t *testing.T) {
	entrada := [][]int{{1, 5}, {4, 20}}
	saidaEsperada := [][]int{{1, 20}}
	saida := OrdenaSlice(entrada)

	assert.Equal(t, saidaEsperada, saida)
}

func TestOrdenaSlice_ComNValores(t *testing.T) {
	entrada := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	saidaEsperada := [][]int{{1, 6}, {8, 10}, {15, 18}}
	saida := OrdenaSlice(entrada)

	assert.Equal(t, saidaEsperada, saida)
}
