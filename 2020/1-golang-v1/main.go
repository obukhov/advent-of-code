package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed getting working dir: %v", err)
	}
	numbers := readFile(wd + "/input.txt")

	task1(numbers)
	task2(numbers)
}

func task1(numbers []int) {
	// Naive approach with deduplication
	for nIndex, n := range numbers {
		for mIndex, m := range numbers[nIndex+1:] {
			if n+m == 2020 {
				log.Printf(
					"[%d:%d, %d:%d] %d * %d = %d",
					nIndex, numbers[nIndex],
					nIndex+mIndex+1, numbers[nIndex+mIndex+1],
					n, m, n*m,
				)
			}
		}
	}
}

func task2(numbers []int) {
	// Naive approach with deduplication
	for nIndex, n := range numbers {
		for mIndex, m := range numbers[nIndex+1:] {
			for rIndex, r := range numbers[nIndex+mIndex+1:] {
				if n+m+r == 2020 {
					log.Printf(
						"[%d:%d, %d:%d, %d:%d] %d * %d * %d = %d",
						nIndex, numbers[nIndex],
						nIndex+mIndex+1, numbers[nIndex+mIndex+1],
						nIndex+mIndex+rIndex+1, numbers[nIndex+mIndex+rIndex+1],
						n, m, r, n*m*r,
					)
				}
			}
		}
	}
}

// taken from https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
func readFile(name string) (numbers []int) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Error converting number '%s': %v", scanner.Text(), err)
		}
		numbers = append(numbers, n)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
