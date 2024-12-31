package main

//func antinodes(data [][]string) int {
//	c := 0
//	arr := alparr()
//	antennas := make([][][]int, len(arr))
//
//	for i, char := range arr {
//		for k, row := range data {
//			for n, val := range row {
//				if val == char {
//					antennas[i] = append(antennas[i], []int{k, n})
//				}
//			}
//		}
//	}
//
//	for i, row := range antennas {
//		for n, _ := range row {
//			for k, _ := range row {
//				vec1 := antennas[i][k][0] - antennas[i][n][0]
//				vec2 := antennas[i][k][1] - antennas[i][n][1]
//				if vec1 == 0 && vec2 == 0 {
//					c--
//				}
//				pos1 := antennas[i][k][0] + vec1
//				pos2 := antennas[i][k][1] + vec2
//				b, a := createantinode(arr, data, pos1, pos2)
//				if b {
//					data[pos1][pos2] = "#"
//				}
//			}
//
//		}
//	}
//	for i := range data {
//		fmt.Println(data[i])
//	}
//	return calcresult(data, c)
//}

func createantinode(arr []string, data [][]string, i, j int) (bool, []int) {
	if i < 0 || i >= len(data) {
		return false, nil
	}
	if j < 0 || j >= len(data[0]) {
		return false, nil
	}
	for k := 0; k < len(arr); k++ {
		if data[i][j] == arr[k] {
			return false, []int{i, j}
		}

	}
	return true, nil

}
