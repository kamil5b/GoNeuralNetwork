package GoNeuralNetwork

import (
	"errors"
	"fmt"
)

func (N NeuralNetwork) Show() {
	fmt.Println("Layers :", len(N.Layers))
	for i, n := range N.Layers {
		fmt.Println("=============Layer ", i+1)
		n.Show()
	}
	fmt.Println("===================")
}

func (N NeuralNetwork) PredictSingle(input []float64) int {

	n_classes := N.Layers[len(N.Layers)-1].N_Neuron
	classes := make([]float64, n_classes)
	for i := range classes {
		classes[i] = float64(i) / float64(n_classes-1)
	}
	for _, n := range N.Layers {
		input = n.calculate(input)
	}
	output := input
	if output[0] != output[1] {

		fmt.Println(output)
	}
	_, _, _, idx := findMinAndMaxFloat64(output)
	return idx
}

func (N NeuralNetwork) Predict(inputs [][]float64) []int {
	predict := make([]int, len(inputs))
	for i, input := range inputs {
		predict[i] = N.PredictSingle(input)
	}
	return predict
}

func (N NeuralNetwork) calculateCost(input, expected []float64) float64 {
	for _, n := range N.Layers {
		input = n.calculate(input)
	}
	//fmt.Println(input)
	return calculateCost(input, expected)

}

func (N *NeuralNetwork) backwardPropagate(expected float64) {
	for i := len((*N).Layers) - 1; i >= 0; i-- {
		layer := ((*N).Layers)[i]
		//error per neuron per layer
		arrError := make([]float64, 0)
		if i != len((*N).Layers)-1 {
			for j := range layer.Neurons {
				err := 0.0
				nextLayer := ((*N).Layers)[i+1]
				for _, neuron := range nextLayer.Neurons {
					//cost/bias derivative
					tmp := ((*neuron).Weight[j] * (*neuron).Delta)
					err += float64(int(tmp*10000)) / 10000
				}
				//err = float64(int(err*10000)) / 10000
				arrError = append(arrError, err)
			}
		} else {
			for _, neuron := range layer.Neurons {
				arrError = append(arrError, (*neuron).tmp_output-expected)
			}
		}
		layer.calculateDelta(arrError)
	}
}

func (N *NeuralNetwork) updateWeight(inputs []float64, lrate float64) {
	tmp_input := inputs
	for i, layer := range (*N).Layers {
		//fmt.Println("updating layer", i+1)
		if i != 0 {
			new_inputs := make([]float64, ((*N).Layers)[i-1].N_Neuron)
			//fmt.Println("neuron:", ((*N).Layers)[i-1].N_Neuron)
			for j, neuron := range ((*N).Layers)[i-1].Neurons {
				new_inputs[j] = (*neuron).tmp_output
			}
			////fmt.Println(len(new_inputs))
			tmp_input = new_inputs
		}

		//fmt.Println(len(tmp_input))
		layer.updateLayer(lrate, tmp_input)
	}
}

func (N *NeuralNetwork) Training(inputs [][]float64, expected []int, lrate float64, epochs int) float64 {
	lastCost := 0.0
	unique := calculateUnique(expected)
	normalizeTarget := make([]float64, len(expected))
	for i := range normalizeTarget {
		normalizeTarget[i] = float64(expected[i]) / float64(unique)
	}
	for e := 0; e < epochs; e++ {
		sumCost := 0.0
		for i, input := range inputs {
			cost := N.calculateCost(input, normalizeTarget)
			sumCost += cost
			N.backwardPropagate(normalizeTarget[i])
			N.updateWeight(input, lrate)
		}
		lastCost = sumCost / float64(len(inputs))
		fmt.Printf("> Epoch=%d, lrate=%.3f, error=%.3f\n", e, lrate, lastCost)
	}
	return lastCost
}

func (N *NeuralNetwork) AddLayer(L Layer) error {
	if len((*N).Layers) > 0 {
		lastLayer := ((*N).Layers)[len((*N).Layers)-1]
		if lastLayer.N_Neuron != L.N_Inputs {
			return errors.New("n input invalid")
		}
	}
	(*N).Layers = append((*N).Layers, &L)
	return nil
}
