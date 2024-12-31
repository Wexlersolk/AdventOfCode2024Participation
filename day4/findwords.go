package main

func exists(matrix [][]string, i int, j int, letter string) bool {
	if i < 0 || i >= len(matrix) {
		return false
	}
	if j < 0 || j >= len(matrix[i]) {
		return false
	}
	return matrix[i][j] == letter
}

func findWords(matrix [][]string) int {
	var count int
	vectors := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for i, row := range matrix {
		for j, value := range row {
			if value == "X" {
				for k, _ := range vectors {
					m, n := i, j
					m += vectors[k][0]
					n += vectors[k][1]

					if exists(matrix, m, n, "M") {
						m += vectors[k][0]
						n += vectors[k][1]

						if exists(matrix, m, n, "A") {
							m += vectors[k][0]
							n += vectors[k][1]

							if exists(matrix, m, n, "S") {
								count++
								
							}
						}
					}
				}
			}

		}
	}
	return count
}
