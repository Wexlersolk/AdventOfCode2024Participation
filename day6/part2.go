package main

import "fmt"

func deepcopy(data [][]string) [][]string {
	copydata := make([][]string, len(data))
	for i := range data {
		copydata[i] = make([]string, len(data[i]))
		copy(copydata[i], data[i])
	}
	return copydata
}

func findguardintime(data [][]string) (int, int) {
	for i := range data {
		for j := range data[i] {
			if data[i][j] == "^" {
				return i, j
			}
		}
	}
	fmt.Printf("Found guard atn")
	return 0, 0
}

func moveguardloop(data [][]string) int {
	vectors := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	m, n := findguardintime(data)
	v := 0
	counter := 0
	for m < len(data) && n < len(data[m]) && counter <= 16000 {
		m, n = findguardintime(data)
		if m+vectors[v][0] >= len(data) || n+vectors[v][1] >= len(data[m]) {
			data[m][n] = "X"
			return 0
		}
		if m+vectors[v][0] < 0 || n+vectors[v][1] < 0 {
			data[m][n] = "X"
			return 0
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
		counter++

	}
	fmt.Println("weak and powerless")
	return 1
}

func timeloop(data [][]string) int {
	count := 0
	for i := range data {
		for j := range data[i] {
			if data[i][j] == "." {
				data[i][j] = "#"
				copydata := deepcopy(data)
				count += moveguardloop(copydata)
				data[i][j] = "."
			}
		}
	}
	return count
}
