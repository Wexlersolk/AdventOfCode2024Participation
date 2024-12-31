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
	var rules [][]int
	var arrays [][]int

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	isRulesSection := true

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			isRulesSection = false
			continue
		}

		parts := strings.Split(line, "|")
		if len(parts) == 1 {

			parts = strings.Split(line, ",")
		}

		var row []int
		for _, part := range parts {
			num, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				log.Fatalf("Error converting %q to integer: %v", part, err)
			}
			row = append(row, num)
		}

		if isRulesSection {
			rules = append(rules, row)
		} else {
			arrays = append(arrays, row)
		}
	}

	fmt.Println(allManuals(rules, arrays))

}

