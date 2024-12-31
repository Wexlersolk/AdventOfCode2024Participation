package quicksort

import (
	"math/rand"
)

func ChoosePivot(length int) int {
	return rand.Intn(length)
}

func Swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func SeparateArray(array []int, pivotPosition int) ([]int, []int) {
	return array[:pivotPosition], array[pivotPosition:]
}

func QuickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	pivotPosition := ChoosePivot(len(array))
	pivot := array[pivotPosition]

	Swap(array, pivotPosition, len(array)-1)

	i := 0
	for j := 0; j < len(array)-1; j++ {
		if array[j] <= pivot {
			Swap(array, i, j)
			i++
		}
	}

	Swap(array, i, len(array)-1)

	left := QuickSort(array[:i])
	right := QuickSort(array[i+1:])

	return append(append(left, pivot), right...)
}
