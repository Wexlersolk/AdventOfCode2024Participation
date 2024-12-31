package main

func checkManual(rules [][]int, arrays []int) ([]int, bool) {
	corrupted := false
	for i := 0; i < len(arrays); i++ {
		for j := 0; j < len(rules); j++ {
			if arrays[i] == rules[j][1] {
				for m := i; m < len(arrays); m++ {
					if arrays[m] == rules[j][0] {
						arrays[i], arrays[m] = arrays[m], arrays[i]
						corrupted = true
					}
				}
			}
		}

	}
	return arrays, corrupted
}

func allManuals(rules [][]int, arrays [][]int) int {
	count := 0
	corrupted := true
	counter := 0
	for i := 0; i < len(arrays); i++ {
		corrupted = true
		counter = 0
		for corrupted {
			arrays[i], corrupted = checkManual(rules, arrays[i])
			counter++
		}
		if counter > 1 {
			count += arrays[i][(len(arrays[i])-1)/2]
		}
	}
	return count
}
