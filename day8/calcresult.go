package main

func calcresult(data [][]string, grabba [][]int) int {
	count := 0

	for _, coords := range grabba {
		if len(coords) != 2 {
			continue
		}
		i, j := coords[0], coords[1]
		if i >= 0 && i < len(data) && j >= 0 && j < len(data[i]) {
			data[i][j] = "#"
		}
	}

	for _, row := range data {
		for _, val := range row {
			if val == "#" {
				count++
			}
		}
	}

	return count
}

