package cmd

import (
	"github.com/obukhov/advent-of-code/2020/golang/lib"
	"github.com/spf13/cobra"
	"log"
	"os"
	"sync"
	"strings"
)

func init() {
	rootCmd.AddCommand(day6cmd)
}

var day6cmd = &cobra.Command{
	Use:   "day6",
	Short: "Day 6 tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day6()
	},
}



func day6() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed getting working dir: %v", err)
	}

	task1input := make(chan string, 100)
	task2input := make(chan string, 100)

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go day6task1(task1input, wg)
	go day6task2(task2input, wg)

	go func() {
		lib.ReadFile(
			wd+"/input/day6.txt",
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

func day6task1(input chan string, wg *sync.WaitGroup) {
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

func day6task2(input chan string, wg *sync.WaitGroup) {
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
