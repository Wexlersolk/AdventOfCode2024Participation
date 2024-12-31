package main

import (
	"bufio"
	"day1/leveldumpener"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("levels.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var safeLevels int
	var currentLevel []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		for _, part := range parts {

			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting to integer:", part)
				continue
			}

			currentLevel = append(currentLevel, num)
		}
		if leveldumpener.IsSafe(currentLevel) {
			safeLevels++
		}
		fmt.Println(currentLevel)
		currentLevel = []int{}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(safeLevels)
}
