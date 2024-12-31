package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var matrix [][]string
	var count int

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}
		matrix = append(matrix, row)
	}
	count += findXmas1(matrix)
	count += findXmas2(matrix)
	count += findXmas3(matrix)
	count += findXmas4(matrix)
	fmt.Println(count)

}
