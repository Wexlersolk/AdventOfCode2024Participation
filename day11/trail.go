package main

func trailscores(data [][]int) int {
	count := 0

	vectors := [][]int{
		{-1, 0},
		{0, -1}, {0, 1},
		{1, 0},
	}
	for i := range data {
		for j := range data[i] {
			if data[i][j] == 0 {
				count += findtrail(data, i, j, vectors)
			}
		}
	}
	return count
}

func findtrail(data [][]int, i, j int, vectors [][]int) int {
	var trails int
	var queue [][]int
	//	var bannedpeaks [][]int
	queue = append(queue, []int{i, j})
	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]

		for _, vector := range vectors {
			newX := current[0] + vector[0]
			newY := current[1] + vector[1]

			if newX >= 0 && newX < len(data) && newY >= 0 && newY < len(data[0]) {

				if data[newX][newY] == data[current[0]][current[1]]+1 {
					//				banned := false
					if data[newX][newY] == 9 {
						//	for m := range bannedpeaks {
						//		if newX == bannedpeaks[m][0] && newY == bannedpeaks[m][1] {
						//			banned = true
						//		}
						//	}
						//
						//	if !banned {
						//		bannedpeaks = append(bannedpeaks, []int{newX, newY})
						//	}
						trails++
						continue
					}
					queue = append(queue, []int{newX, newY})
				}
			}
		}
	}
	return trails
}
