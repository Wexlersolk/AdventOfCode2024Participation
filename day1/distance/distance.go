package distance

func AbsoluteValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func CalculateDistance(array1 []int, array2 []int) int {
	var distance int

	if len(array1) != len(array2) {
		return -1
	}

	for i := 0; i < len(array1); i++ {
		distance += AbsoluteValue(array1[i] - array2[i])
	}
	return distance
}

