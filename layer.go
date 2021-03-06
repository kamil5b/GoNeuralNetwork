package GoNeuralNetwork

import (
	"fmt"
)

func (L Layer) Show() {
	fmt.Println("N_Input :", L.N_Inputs)
	fmt.Println("N_Output :", L.N_Neuron, len(L.Neurons))
	for i, n := range L.Neurons {
		fmt.Println("====Neuron ", i+1)
		n.Show()
	}
}

func (L Layer) calculate(inputs []float64) []float64 {
	outputs := make([]float64, len(L.Neurons))
	for i, neuron := range L.Neurons {
		outputs[i] = (*neuron).calculate(inputs, L.ActivationFunc)
	}
	//fmt.Println("===OUTPUT", outputs)
	return outputs
}

func (layer *Layer) calculateDelta(err []float64) {
	for i, neuron := range (*layer).Neurons {
		td := layer.DerivativeFunc((*neuron).tmp_sum)
		tmp := err[i] * td
		(*neuron).Delta = float64(int(tmp*10000)) / 10000
	}
}

func (layer *Layer) updateLayer(lrate float64, inputs []float64) {

	for _, neuron := range (*layer).Neurons {
		//fmt.Println("Weight", (*neuron).Weight, layer.N_Inputs)
		//fmt.Println("Inputs", inputs)
		//fmt.Println("===================")
		for j := range inputs {
			//INI SETIAP WEIGHT
			tmp := lrate * neuron.Delta * inputs[j]
			neuron.Weight[j] -= float64(int(tmp*10000)) / 10000
		}
		//INI BIAS
		tmp := lrate * neuron.Delta
		neuron.Bias -= float64(int(tmp*10000)) / 10000
	}
}
