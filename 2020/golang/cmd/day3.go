package cmd

import (
	"github.com/obukhov/advent-of-code/2020/golang/lib"
	"github.com/spf13/cobra"
	"log"
	"os"
	"sync"
)

func init() {
	rootCmd.AddCommand(day3cmd)
}

var day3cmd = &cobra.Command{
	Use:   "day3",
	Short: "Day 3 tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day3()
	},
}

const (
	TREE = '#'
)

type slope struct {
	xStep int
	yStep int
}

func day3() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed getting working dir: %v", err)
	}

	task1input := make(chan string, 100)
	task2input := make(chan string, 100)

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go day3task1(task1input, wg)
	go day3task2(task2input, wg)

	go func() {
		lib.ReadFile(
			wd+"/input/day3.txt",
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

func day3task1(input chan string, wg *sync.WaitGroup) {
	var (
		pos int
		total int
	)
	for line := range input {
		if line[pos] == TREE {
			total++
		}
		pos = (pos+3) % len(line)
	}

	log.Printf("Total trees: %d", total)
	wg.Done()
}


func day3task2(input chan string, wg *sync.WaitGroup) {
	var (
		lineNum int
		slopes = []slope{
			{1, 1},
			{3, 1},
			{5, 1},
			{7, 1},
			{1, 2},
		}
	)
	slopesCount := len(slopes)
	positions := make([]int, slopesCount, slopesCount)
	totals := make([]int, slopesCount, slopesCount)

	for line := range input {
		for nSlope, slope := range slopes {
			if lineNum % slope.yStep == 0 {
				if line[positions[nSlope]] == TREE {
					totals[nSlope]++
				}

				positions[nSlope] = (positions[nSlope] + slope.xStep) % len(line)
			}
		}

		lineNum++
	}

	totalMultiplication := 1
	for _, t := range totals {
		totalMultiplication *= t
	}

	log.Printf("Total trees on each slope: %v, multiplication: %v", totals, totalMultiplication)
	wg.Done()
}