package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"

)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed getting working dir: %v", err)
	}

	records := readFile(wd + "/input.txt")

	task1(records)
	task2(records)
}

func task1(records []record) {
	var total int
	for _, record := range records {
		var (
			cnt     int
		)
		for _, b := range record.password {
			if b == record.symbol {
				cnt++
			}
		}

		if record.min <= cnt && cnt <= record.max {
			total++
		}
	}

	log.Printf("Total number of correct passwords 1: %d", total)
}

func task2(records []record) {
	var total int
	for _, record := range records {
		pos1match := rune(record.password[record.min-1]) == record.symbol
		pos2match := rune(record.password[record.max-1]) == record.symbol
		if pos1match != pos2match {
			total++
		}
	}

	log.Printf("Total number of correct passwords 2: %d", total)
}

type record struct {
	min      int
	max      int
	symbol   rune
	password string
}

// taken from https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
func readFile(name string) (records []record) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	exp, err := regexp.Compile("^(\\d+)-(\\d+)\\s(\\S): (\\S+)$")
	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		matches := exp.FindAllStringSubmatch(scanner.Text(), 4)
		if len(matches) < 1 || len(matches[0]) != 5 {
			log.Fatalf("Error matching string '%s'", scanner.Text())
		}

		min, _ := strconv.Atoi(matches[0][1])
		max, _ := strconv.Atoi(matches[0][2])

		records = append(records, record{
			min,
			max,
			rune(matches[0][3][0]),
			matches[0][4],
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
