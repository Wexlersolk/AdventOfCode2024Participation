package main

import (
	"bufio"
	"day1/distance"
	"day1/quicksort"
	"day1/similarity"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("locations.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var firstArray []int
	var secondArray []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		firstNum, err1 := strconv.Atoi(parts[0])
		secondNum, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error converting numbers in line:", line)
			continue
		}
		firstArray = append(firstArray, firstNum)
		secondArray = append(secondArray, secondNum)
	}

	firstArray = quicksort.QuickSort(firstArray)
	secondArray = quicksort.QuickSort(secondArray)

	distance := distance.CalculateDistance(firstArray, secondArray)
	similarity := similarity.CalculateSimilarity(firstArray, secondArray)

	fmt.Println(distance)
	fmt.Println(similarity)

}
