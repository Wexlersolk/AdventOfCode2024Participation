package main

func findguard(data [][]string) (int, int) {
	for i := range data {
		for j := range data[i] {
			if data[i][j] == "^" {
				return i, j
			}
		}
	}
	return 0, 0
}

func moveguard(data [][]string) [][]string {
	vectors := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	m, n := findguard(data)
	v := 0
	for i := 0; i < 16900; i++ {
		for m < len(data) && n < len(data[m]) {
			m, n = findguard(data)
			if m+vectors[v][0] >= len(data) || n+vectors[v][1] >= len(data[m]) {
				data[m][n] = "X"
				return data
			}
			if m+vectors[v][0] < 0 || n+vectors[v][1] < 0 {
				data[m][n] = "X"
				return data
			}
			if data[m+vectors[v][0]][n+vectors[v][1]] != "#" {
				data[m][n] = "X"
				data[m+vectors[v][0]][n+vectors[v][1]] = "^"
			} else {
				v++
				if v == 4 {
					v = 0
				}
			}

		}
	}
	return data
}
