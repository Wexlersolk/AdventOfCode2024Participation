package similarity

func numberAppearance(array2 []int, number int) int {
	j := 0
	for i := 0; i < len(array2); i++ {
		if number == array2[i] {
			j++
		}
	}
	return j
}

func CalculateSimilarity(array1 []int, array2 []int) int {
	similarity := 0
	for i := 0; i < len(array1); i++ {
		similarity += array1[i] * numberAppearance(array2, array1[i])
	}
	return similarity
}
