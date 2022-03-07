package GoNeuralNetwork

import "math"

func sigmoid(input float64) float64 {
	inp := float64(int(input*10000)) / 10000
	return 1.0 / (1.0 + math.Exp(-inp))
}

func elU(input float64) float64 {
	inp := float64(int(input*10000)) / 10000
	if inp > 0 {
		return inp
	}
	return 1.0 * (math.Exp(inp) - 1)
}

func relU(input float64) float64 {
	inp := float64(int(input*10000)) / 10000
	if inp > 0 {
		return inp
	}
	return 0
}

func leakyRelU(input float64) float64 {
	inp := float64(int(input*10000)) / 10000
	if input > 0.1*inp {
		return inp
	}
	return 0.1 * inp
}

func tanh(input float64) float64 {
	inp := float64(int(input*10000)) / 10000
	return math.Tanh(inp)
}

func linear(input float64) float64 {
	inp := float64(int(input*10000)) / 10000
	return inp
}

func derivativeSigmoid(input float64) float64 {
	//σ(x)⋅(1−σ(x))
	inp := float64(int(input*10000)) / 10000
	return sigmoid(inp) * (1 - sigmoid(inp))
}

func derivativeElu(input float64) float64 {
	inp := float64(int(input*10000)) / 10000
	if input > 0 {
		return 1
	}
	return 1.0 * (math.Exp(inp))
}

func derivativeRelU(input float64) float64 {
	inp := float64(int(input*10000)) / 10000
	if inp < 0 {
		return 0
	}
	return 1
}

func derivativeLeakyRelU(input float64) float64 {
	inp := float64(int(input*10000)) / 10000
	if inp < 0 {
		return 0.01
	}
	return 1
}

func derivativeTanh(input float64) float64 {
	inp := float64(int(input*10000)) / 10000
	//1 - tanh^2(x)
	return 1 - (math.Tanh(inp) * math.Tanh(inp))
}

func derivativeLinear(input float64) float64 {
	return 1.0
}

//============USABLE=============
func Sigmoid() (func(input float64) float64, func(input float64) float64) {
	return sigmoid, derivativeSigmoid
}

func ElU() (func(input float64) float64, func(input float64) float64) {
	return elU, derivativeElu
}

func RelU() (func(input float64) float64, func(input float64) float64) {
	return relU, derivativeRelU
}

func LeakyRelU() (func(input float64) float64, func(input float64) float64) {
	return leakyRelU, derivativeLeakyRelU
}

func Tanh() (func(input float64) float64, func(input float64) float64) {
	return tanh, derivativeTanh
}

func Linear() (func(input float64) float64, func(input float64) float64) {
	return linear, derivativeLinear
}
