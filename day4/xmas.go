package main

func findXmas1(matrix [][]string) int {
	var count int
	positions := [][]int{
		{0, 2}, {1, 1}, {2, 2}, {2, 0},
	}
	letters := []string{"M", "A", "S", "S"}

	for i, row := range matrix {
		for j, value := range row {
			if value == "M" {
				counter := 0
				for k, _ := range positions {
					m, n := i, j
					m += positions[k][0]
					n += positions[k][1]

					if exists(matrix, m, n, letters[counter]) {
						counter++

					}
					if counter == 4 {
						count++
					}
				}
			}
		}

	}
	return count
}

func findXmas2(matrix [][]string) int {
	var count int
	positions := [][]int{
		{0, 2}, {1, 1}, {2, 2}, {2, 0},
	}
	letters := []string{"S", "A", "M", "M"}

	for i, row := range matrix {
		for j, value := range row {
			if value == "S" {
				counter := 0
				for k, _ := range positions {
					m, n := i, j
					m += positions[k][0]
					n += positions[k][1]

					if exists(matrix, m, n, letters[counter]) {
						counter++

					}
					if counter == 4 {
						count++
					}
				}
			}
		}

	}
	return count
}

func findXmas3(matrix [][]string) int {
	var count int
	positions := [][]int{
		{0, 2}, {1, 1}, {2, 2}, {2, 0},
	}
	letters := []string{"M", "A", "M", "S"}

	for i, row := range matrix {
		for j, value := range row {
			if value == "S" {
				counter := 0
				for k, _ := range positions {
					m, n := i, j
					m += positions[k][0]
					n += positions[k][1]

					if exists(matrix, m, n, letters[counter]) {
						counter++

					}
					if counter == 4 {
						count++
					}
				}
			}
		}

	}
	return count
}

func findXmas4(matrix [][]string) int {
	var count int
	positions := [][]int{
		{0, 2}, {1, 1}, {2, 2}, {2, 0},
	}
	letters := []string{"S", "A", "S", "M"}

	for i, row := range matrix {
		for j, value := range row {
			if value == "M" {
				counter := 0
				for k, _ := range positions {
					m, n := i, j
					m += positions[k][0]
					n += positions[k][1]

					if exists(matrix, m, n, letters[counter]) {
						counter++

					}
					if counter == 4 {
						count++
					}
				}
			}
		}

	}
	return count
}
