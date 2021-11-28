package ordena

import "fmt"

// in: [1,3], [2,6], [8,10], [15,18]
// out: [1, 6],[8,10],[15,18]
// in: [1,3], [8,10], [15, 18]
// out: [1,3], [8,10], [15, 18]
// in: [1,3], [2,6], [6, 50]
// out: [1,50]

func SomaDeMultiplos(N int) int {
	soma := 0
	for i := 3; i <= N; i++ {
		if i%3 == 0 || i%5 == 0 {
			soma += i
		}
	}
	return soma
}

func OrdenaSlice(matriz [][]int) [][]int {
	var item1 int
	var item2 int
	var matrizRetorno [][]int
	firstEntry := true
	// iteracao de primeiro nivel
	for i, m := range matriz {
		//fmt.Println(m)
		if firstEntry {
			item2 = m[1]
			item1 = m[0]
			firstEntry = false
			continue
		}

		fmt.Println(item1, item2)
		if m[0] <= matriz[i-1][1] {
			item2 = m[1]
		} else {
			matrizRetorno = append(matrizRetorno, []int{item1, item2})
			item1 = m[0]
			item2 = m[1]
		}
	}
	matrizRetorno = append(matrizRetorno, []int{item1, item2})
	return matrizRetorno
}
