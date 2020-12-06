package main

import (
	"github.com/obukhov/advent-of-code/common"
	"log"
	"os"
	"regexp"
	"strconv"
	"sync"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed getting working dir: %v", err)
	}

	exp, err := regexp.Compile("^(\\d+)-(\\d+)\\s(\\S): (\\S+)$")
	if err != nil {
		log.Fatal(err)
	}

	task1input := make(chan record, 100)
	task2input := make(chan record, 100)

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go task1(task1input, wg)
	go task2(task2input, wg)
	go func() {
		common.ReadFile(
			wd + "/../2-golang-v1/input.txt",
			func (line string) {
				matches := exp.FindAllStringSubmatch(line, 4)
				if len(matches) < 1 || len(matches[0]) != 5 {
					log.Fatalf("Error matching string '%s'", line)
				}

				min, _ := strconv.Atoi(matches[0][1])
				max, _ := strconv.Atoi(matches[0][2])

				r := record{
					min,
					max,
					rune(matches[0][3][0]),
					matches[0][4],
				}

				task1input <- r
				task2input <- r

			},
			wg,
		)

		close(task1input)
		close(task2input)
	}()

	wg.Wait()
}

func task1(input chan record, wg *sync.WaitGroup) {
	defer wg.Done()

	var total int
	for record := range input {
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

func task2(input chan record, wg *sync.WaitGroup) {
	defer wg.Done()

	var total int
	for record := range input {
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