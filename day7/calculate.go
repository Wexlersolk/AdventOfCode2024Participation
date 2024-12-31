package main

func calculate(data [][]int) int {
	count := 0
	for i := range data {
		numbers := data[i][1:]
		operatorsCombos := generateOperators(len(numbers))
		matchFound := false // Ensure each target is only added once

		for _, operators := range operatorsCombos {
			value := numbers[0]
			for j := 0; j < len(operators); j++ {
				switch operators[j] {
				case "+":
					value += numbers[j+1]
				case "*":
					value *= numbers[j+1]
				}
			}

			if value == data[i][0] {
				if !matchFound { // Add only once for the first match
					count += value
					matchFound = true
				}
			}
		}
	}
	return count
}

func generateOperators(n int) [][]string {
	var result [][]string

	numOperators := 1 << (n - 1)
	for i := 0; i < numOperators; i++ {
		var operators []string
		for j := 0; j < n-1; j++ {
			if (i & (1 << j)) != 0 {
				operators = append(operators, "*")
			} else {
				operators = append(operators, "+")
			}
		}
		result = append(result, operators)
	}
	return result

}
