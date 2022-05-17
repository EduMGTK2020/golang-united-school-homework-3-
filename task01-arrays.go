package homework

func average(input [15]float32) (result float32) {
	var sum float32
	var i int8
	for _, in := range input {
		if in > 0.0 {
			sum = sum + in
			i++
		}
	}
	result = sum / float32(i)
	return
}
