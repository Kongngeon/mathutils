package github.com/Kongngeon/mathutils

func Average(input []float64) float64 {
	sum := 0.0
	for _, i := range input {
		sum += i
	}

	avg := sum / float64(len(input))
	return avg
}
