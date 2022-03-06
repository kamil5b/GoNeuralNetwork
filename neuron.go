package GoNeuralNetwork

import "fmt"

func (N Neuron) Show() {
	fmt.Println("N_Input :", len(N.Weight))
	fmt.Println("Weight :", N.Weight)
	fmt.Println("Bias :", N.Bias)
	fmt.Println("Delta :", N.Delta)
}

func (N *Neuron) calculate(inputs []float64, activationFunc func(float64) float64) float64 {
	sum := 0.0
	//fmt.Println("Weight", (N.Weight))
	//fmt.Println("Inputs", inputs)
	//fmt.Println("===================")
	for i := 0; i < len(N.Weight); i++ {
		x := N.Weight[i] * inputs[i]
		sum += float64(int(x*10000)) / 10000
	}
	sum += N.Bias
	N.tmp_sum = sum
	N.tmp_output = activationFunc(sum)
	N.tmp_input = inputs
	return N.tmp_output
}
