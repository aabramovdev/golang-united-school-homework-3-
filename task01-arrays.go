package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var sum float32
	var length int
	for i := 0; i < 15; i++ {
		if input[i] == 0 {
			break
		}
		sum += input[i]
		length += 1
	}
	result = sum / float32(length)
	return
}
