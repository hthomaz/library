package calculadora

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Input struct {
	Expression string `json:"expression"`
}

type Output struct {
	Answer float64 `json:"answer"`
}

func calculadoraHandler(w http.ResponseWriter, r *http.Request) {
	buffer, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Header().Add("status", "400")
		return
	}

	input := &Input{}
	err = json.Unmarshal(buffer, input)

	if err != nil {
		w.Header().Add("status", "400")
		return
	}

	resposta, err := Calcula(input.Expression)
	if err != nil {
		w.Header().Add("status", "400")
		return
	}

	output := &Output{
		Answer: resposta,
	}

	bufferOuput, _ := json.Marshal(output)

	w.Write(bufferOuput)

	return
}

func calculadoraMultiplasExpressoesHandler(w http.ResponseWriter, r *http.Request) {
	buffer, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Header().Add("status", "400")
		fmt.Println(err)
		return
	}

	var inputs []Input
	err = json.Unmarshal(buffer, &inputs)

	if err != nil {
		w.Header().Add("status", "400")
		fmt.Println(err)
		return
	}

	var outputs []*Output
	for _, input := range inputs {
		reposta, _ := Calcula(input.Expression)
		output := &Output{
			Answer: reposta,
		}
		outputs = append(outputs, output)
	}
	fmt.Println(outputs)

	bufferOuput, _ := json.Marshal(outputs)

	w.Write(bufferOuput)

	return
}

func Calculadora() {
	http.HandleFunc("/calculadora", calculadoraHandler)
	http.HandleFunc("/calculadoraMultiplasExpressoes", calculadoraMultiplasExpressoesHandler)
	http.ListenAndServe(":3000", nil)
}
