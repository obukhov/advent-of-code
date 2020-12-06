package main

import (
	"github.com/obukhov/advent-of-code/common"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed getting working dir: %v", err)
	}

	task1input := make(chan string, 100)
	task2input := make(chan string, 100)

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go task1(task1input, wg)
	go task2(task2input, wg)

	go func() {
		common.ReadFile(
			wd + "/input.txt",
			func (line string) {
				task1input <- line
				task2input <- line
			},
			wg,
		)

		close(task1input)
		close(task2input)
	}()

	wg.Wait()
}

func task1(input chan string, wg *sync.WaitGroup) {
	uniqueAnswers := make(map[rune]bool)
	answersCount := make([]int, 0)
	totalCount := 0

	for answers := range input {
		if answers == "" {
			answersCount = append(answersCount, len(uniqueAnswers))
			totalCount += len(uniqueAnswers)
			uniqueAnswers = make(map[rune]bool)
		} else {
			for _, answer := range answers {
				uniqueAnswers[answer] = true
			}
		}
	}

	answersCount = append(answersCount, len(uniqueAnswers))
	totalCount += len(uniqueAnswers)

	log.Printf("Group anyone counts %v, total: %d", answersCount, totalCount)
	wg.Done()
}

func task2(input chan string, wg *sync.WaitGroup) {
	uniqueAnswers := make(map[rune]bool)
	answersCount := make([]int, 0)
	totalCount := 0
	isFirstInGroup := true

	for answers := range input {
		if answers == "" {
			answersCount = append(answersCount, len(uniqueAnswers))
			totalCount += len(uniqueAnswers)

			uniqueAnswers = make(map[rune]bool)
			isFirstInGroup = true
		} else {
			if isFirstInGroup {
				for _, answer := range answers {
					uniqueAnswers[answer] = true
				}
				isFirstInGroup = false
			} else {
				for existingAnswer, _ := range uniqueAnswers {
					if !strings.Contains(answers, string(existingAnswer)) {
						delete(uniqueAnswers, existingAnswer)
					}
				}
			}
		}
	}

	answersCount = append(answersCount, len(uniqueAnswers))
	totalCount += len(uniqueAnswers)

	log.Printf("Group everyone counts %v, total: %d", answersCount, totalCount)
	wg.Done()
}
