package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

		parts := strings.Split(line, " ")

		key := strings.TrimSuffix(parts[0], ":")
		keyNum, err := strconv.Atoi(key)
		if err != nil {
			log.Printf("Skipping invalid key %q: %v\n", parts[0], err)
			continue
		}

		row := []int{keyNum}

		for _, part := range parts[1:] {
			num, err := strconv.Atoi(part)
			if err != nil {
				log.Printf("Skipping invalid number %q: %v\n", part, err)
				continue
			}
			row = append(row, num)
		}

		data = append(data, row)
	}

	fmt.Println(calculate2(data))

}
