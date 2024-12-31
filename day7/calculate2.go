package main

import (
	"fmt"
	"math"
	"strconv"
)

func calculate2(data [][]int) int {
	count := 0
	for i := range data {
		originalNumbers := data[i][1:]
		operatorsCombos := generateOperators2(len(originalNumbers))
		matchFound := false
		for _, operators := range operatorsCombos {

			numbers := append([]int{}, originalNumbers...)
			operatorsCopy := append([]string{}, operators...)

			value := numbers[0]
			if data[i][0] == 7290 {
				fmt.Println(numbers, operatorsCopy)
			}

			for j := 0; j < len(operatorsCopy); j++ {
				if j+1 < len(numbers) {
					switch operatorsCopy[j] {
					case "+":
						value += numbers[j+1]
					case "*":
						value *= numbers[j+1]
					case "|":
						num1 := strconv.Itoa(value)
						num2 := strconv.Itoa(numbers[j+1])
						result := num1 + num2
						value, _ = strconv.Atoi(result)
					}
				}
			}

			if value == data[i][0] {
				if !matchFound {
					count += value
					matchFound = true
				}
			}
		}
	}
	return count
}

func generateOperators2(n int) [][]string {
	var result [][]string

	numOperators := int(math.Pow(3, float64(n-1)))

	for i := 0; i < numOperators; i++ {
		var operators []string
		value := i
		for j := 0; j < n-1; j++ {
			switch value % 3 {
			case 0:
				operators = append(operators, "+")
			case 1:
				operators = append(operators, "*")
			case 2:
				operators = append(operators, "|")
			}
			value /= 3
		}
		result = append(result, operators)
	}

	return result
}
