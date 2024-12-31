package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data [][]string

	for scanner.Scan() {
		line := scanner.Text()

		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}

		data = append(data, row)
	}
	count := antinodes2(data)
	fmt.Println(count)
}
