package GoNeuralNetwork

import "math/rand"

func findMinAndMaxFloat64(a []float64) (min, max float64, imin, imax int) {
	min, max = a[0], a[0]
	imin, imax = 0, 0
	for i, value := range a {
		if value < min {
			min = value
			imin = i
		}
		if value > max {
			max = value
			imax = i
		}
	}
	return min, max, imin, imax
}

func calculateUnique(target []int) int {
	tmptarget := make(map[int]bool)
	for _, s := range target {
		if _, value := tmptarget[int(s)]; !value {
			tmptarget[int(s)] = true
		}
	}
	return len(tmptarget)
}

func AccuracyScore(predictions []int, actual []int) float64 {
	correct := float64(0)
	for i := range predictions {
		if predictions[i] == actual[i] {
			correct += 1
		}
	}
	return correct / float64(len(predictions))
}

func ConfusionMatrix(predictions []int, actual []int) [][]int {
	n_unique := calculateUnique(actual)
	matrix := make([][]int, n_unique)
	for i := 0; i < n_unique; i++ {
		matrix[i] = make([]int, n_unique)
	}
	for i := range predictions {
		matrix[actual[i]][predictions[i]] += 1
	}
	return matrix
}

func RandFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
