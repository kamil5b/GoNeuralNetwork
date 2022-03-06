package GoNeuralNetwork

import "math"

func softmax(input float64) float64 {
	return 1.0 / (1.0 + math.Exp(-input))
}

/*
func relU(input float64) float64 {
	if input > 0 {
		return input
	}
	return 0
}

func leakyRelU(input float64) float64 {
	if input > 0.1*input {
		return input
	}
	return 0.1 * input
}*/

func tanh(input float64) float64 {
	return math.Tanh(input)
}

func linear(input float64) float64 {
	return input
}

func derivativeSoftMax(input float64) float64 {
	//σ(x)⋅(1−σ(x))
	return softmax(input) * (1 - softmax(input))
}

/*
func derivativeRelU(input float64) float64 {
	if input < 0 {
		return 0
	}
	return 1
}

func derivativeLeakyRelU(input float64) float64 {
	if input < 0 {
		return 0.01
	}
	return 1
}*/

func derivativeTanh(input float64) float64 {
	//1 - tanh^2(x)
	return 1 - (math.Tanh(input) * math.Tanh(input))
}

func derivativeLinear(input float64) float64 {
	return 1.0
}

//============USABLE=============
func SoftMax() (func(input float64) float64, func(input float64) float64) {
	return softmax, derivativeSoftMax
}

/*
func RelU() (func(input float64) float64, func(input float64) float64) {
	return relU, derivativeRelU
}

func LeakyRelU() (func(input float64) float64, func(input float64) float64) {
	return leakyRelU, derivativeLeakyRelU
}*/

func Tanh() (func(input float64) float64, func(input float64) float64) {
	return tanh, derivativeTanh
}

func Linear() (func(input float64) float64, func(input float64) float64) {
	return linear, derivativeLinear
}
