package main

import "strconv"

func calculate(disk []string) []string {
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] != "." {
			for j := 0; j < len(disk); j++ {
				if disk[j] == "." {
					disk[i], disk[j] = disk[j], disk[i]
				}
			}
		}
	}
	disk = disk[1:]
	return disk
}

func count(disk []string) int {
	count := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] != "." {
			v, _ := strconv.Atoi(disk[i])
			count += i * v
		}
	}
	return count
}
