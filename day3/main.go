package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Mul(num1 int, num2 int) int {
	return num1 * num2
}

func main() {
	pattern := `mul\((\d+),\s*(\d+)\)`
	doPattern := `do\(\)`
	dontPattern := `don't\(\)`

	re := regexp.MustCompile(pattern)
	doRe := regexp.MustCompile(doPattern)
	dontRe := regexp.MustCompile(dontPattern)

	total := 0
	enable := true

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileStat, _ := file.Stat()
	if fileStat.Size() == 0 {
		log.Fatal("File is empty")
	}

	data := make([]byte, fileStat.Size())
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	content := string(data)
	lineRe := regexp.MustCompile(`(do\(\)|don't\(\)|mul\(\d+,\s*\d+\))`)
	tokens := lineRe.FindAllString(content, -1)

	for _, token := range tokens {
		if doRe.MatchString(token) {
			enable = true
		} else if dontRe.MatchString(token) {
			enable = false
		} else if re.MatchString(token) {
			match := re.FindStringSubmatch(token)
			num1, err1 := strconv.Atoi(match[1])
			num2, err2 := strconv.Atoi(match[2])
			if err1 != nil || err2 != nil {
				log.Printf("Error converting numbers: %v, %v\n", err1, err2)
				continue
			}
			if enable {
				total += Mul(num1, num2)
			}
		}
	}

	fmt.Println("Total:", total)
}

