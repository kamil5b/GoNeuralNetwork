package GoNeuralNetwork

type Neuron struct {
	Weight []float64 //Weight itu bobot setiap input
	Delta  float64
	Bias   float64
	Active bool

	tmp_input  []float64
	tmp_sum    float64
	tmp_output float64
}
type Layer struct {
	N_Inputs       int
	N_Neuron       int
	Neurons        []*Neuron
	ActivationFunc func(float64) float64
	DerivativeFunc func(float64) float64
	DropOutRate    float64
}
type NeuralNetwork struct {
	Layers   []*Layer
	N_inputs int
}

func CreateNeuron(n_inputs int) Neuron {
	weight := make([]float64, n_inputs)
	for i := 0; i < n_inputs; i++ {
		tmp := RandFloat(-1.0, 1.0)
		weight[i] = float64(int(tmp*10000)) / 10000
	}
	tmp := RandFloat(-1.0, 1.0)
	bias := float64(int(tmp*10000)) / 10000
	return Neuron{
		Weight: weight,
		Bias:   bias,
		Delta:  0,
		Active: true,
	}
}

func CreateLayer(n_inputs, n_neuron int, activationFunc func() (func(input float64) float64, func(input float64) float64)) Layer {
	neurons := make([]*Neuron, n_neuron)
	for i := 0; i < n_neuron; i++ {
		x := CreateNeuron(n_inputs)
		neurons[i] = &x
	}
	act, dev := activationFunc()
	return Layer{
		N_Inputs:       n_inputs,
		N_Neuron:       n_neuron,
		Neurons:        neurons,
		ActivationFunc: act,
		DerivativeFunc: dev,
	}
}

func InitializeNetwork(n_input int) NeuralNetwork {
	return NeuralNetwork{
		make([]*Layer, 0),
		n_input,
	}
}

func calculateCost(predict, expected []float64) float64 {
	cost := 0.0
	//fmt.Println(predict, expected)
	for i := range predict {
		cost += (predict[i] - expected[i]) *
			(predict[i] - expected[i])
	}
	return cost
}
