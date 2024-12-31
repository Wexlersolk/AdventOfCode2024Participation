package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data [][]int

	for scanner.Scan() {
		line := scanner.Text()

		var row []int
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			row = append(row, num)
		}

		data = append(data, row)
	}
	count := trailscores(data)
	fmt.Println(count)
}
