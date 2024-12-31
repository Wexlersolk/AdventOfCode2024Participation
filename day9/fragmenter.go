package main

import (
	"fmt"
	"strconv"
)

func fragment(data []string) []string {
	var disk []string
	id := -1
	for i := range data {
		if i%2 == 0 {
			id++
			v, _ := strconv.Atoi(data[i])
			for m := 0; m < v; m++ {
				disk = append(disk, strconv.Itoa(id))
			}

		} else {
			v, _ := strconv.Atoi(data[i])
			for m := 0; m < v; m++ {
				disk = append(disk, ".")
			}

		}

	}
	fmt.Println(disk)
	return disk
}
