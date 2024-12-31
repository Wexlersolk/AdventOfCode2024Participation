package main

func calculatepositions(data [][]string) int {
	count := 0
	for i := range data {
		for j := range data[i] {
			if data[i][j] == "X" {
				count++
			}
		}
	}
	return count
}
