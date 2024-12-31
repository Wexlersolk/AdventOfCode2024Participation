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

	var data []string

	for scanner.Scan() {
		line := scanner.Text()

		for _, char := range line {
			data = append(data, string(char))
		}
	}
	data = fragment(data)
	count := count(calculate2(data))
	fmt.Println(count)
}
