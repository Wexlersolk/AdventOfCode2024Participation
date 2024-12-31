package main

import "fmt"

func antinodes2(data [][]string) int {
	arr := alparr()
	antennas := make([][][]int, len(arr))
	var grabba [][]int

	for i, char := range arr {
		for k, row := range data {
			for n, val := range row {
				if val == char {
					antennas[i] = append(antennas[i], []int{k, n})
				}
			}
		}
	}

	for i, row := range antennas {
		for n, _ := range row {
			for k, _ := range row {
				vec1 := antennas[i][k][0] - antennas[i][n][0]
				vec2 := antennas[i][k][1] - antennas[i][n][1]
				pos1 := antennas[i][k][0]
				pos2 := antennas[i][k][1]
				for m := 0; m < len(data)*len(data); m++ {
					pos1 += vec1
					pos2 += vec2
					b, a := createantinode(arr, data, pos1, pos2)
					if a != nil {
						grabba = append(grabba, a)
					}
					if b {
						data[pos1][pos2] = "#"
					}
				}
			}

		}
	}
	for i := range data {
		fmt.Println(data[i])
	}
	return calcresult(data, grabba)
}
