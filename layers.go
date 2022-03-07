package GoNeuralNetwork

func (N *NeuralNetwork) AddDenseLayer(n_neuron int) {
	n_input := N.N_inputs
	if len((*N).Layers) > 0 {
		n_input = ((*N).Layers)[len((*N).Layers)-1].N_Neuron
	}
	N.AddLayer(CreateLayer(n_input, n_neuron, Linear))
}

func (N *NeuralNetwork) AddSigmoidLayer(n_neuron int) {
	n_input := N.N_inputs
	if len((*N).Layers) > 0 {
		n_input = ((*N).Layers)[len((*N).Layers)-1].N_Neuron
	}
	N.AddLayer(CreateLayer(n_input, n_neuron, Sigmoid))

}

func (N *NeuralNetwork) AddElULayer(n_neuron int) {
	n_input := N.N_inputs
	if len((*N).Layers) > 0 {
		n_input = ((*N).Layers)[len((*N).Layers)-1].N_Neuron
	}
	N.AddLayer(CreateLayer(n_input, n_neuron, ElU))

}

func (N *NeuralNetwork) AddRelULayer(n_neuron int) {
	n_input := N.N_inputs
	if len((*N).Layers) > 0 {
		n_input = ((*N).Layers)[len((*N).Layers)-1].N_Neuron
	}
	N.AddLayer(CreateLayer(n_input, n_neuron, RelU))

}

func (N *NeuralNetwork) AddLeakyRelULayer(n_neuron int) {
	n_input := N.N_inputs
	if len((*N).Layers) > 0 {
		n_input = ((*N).Layers)[len((*N).Layers)-1].N_Neuron
	}
	N.AddLayer(CreateLayer(n_input, n_neuron, LeakyRelU))

}

func (N *NeuralNetwork) AddTanhLayer(n_neuron int) {
	n_input := N.N_inputs
	if len((*N).Layers) > 0 {
		n_input = ((*N).Layers)[len((*N).Layers)-1].N_Neuron
	}
	N.AddLayer(CreateLayer(n_input, n_neuron, Tanh))

}
