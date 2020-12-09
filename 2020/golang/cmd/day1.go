package cmd

import (
	"github.com/obukhov/advent-of-code/2020/golang/lib"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"sync"
)

func init() {
	rootCmd.AddCommand(day1cmd)
}

var day1cmd = &cobra.Command{
	Use:   "day1",
	Short: "Day 1 tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day1()
	},
}

func day1() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed getting working dir: %v", err)
	}

	task1input := make(chan int, 200)
	task2input := make(chan int, 200)

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go day1task1(task1input, wg)
	go day1task2(task2input, wg)
	go func() {
		lib.ReadFileWg(
			wd+"/input/day1.txt",
			func(line string) {
				n, err := strconv.Atoi(line)
				if err != nil {
					log.Fatalf("Error converting number '%s': %v", line, err)
				}
				task1input <- n
				task2input <- n
			},
			wg,
		)

		close(task1input)
		close(task2input)
	}()

	wg.Wait()
}

func day1task1(input chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	numbers := make([]int, 0)
	for {
		select {
		case m, open:= <-input:
			if !open {
				log.Printf("Task 1 has completed")
				return
			}
			for nIndex, n := range numbers {
				if n+m == 2020 {
					log.Printf(
						"[%d:%d, %d:%d] %d * %d = %d",
						nIndex, numbers[nIndex],
						len(numbers)+1, m,
						n, m, n*m,
					)
				}
			}
			numbers = append(numbers, m)
		}
	}
}

func day1task2(input chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	numbers := make([]int, 0)
	for {
		select {
		case r, open := <-input:
			if !open {
				log.Printf("Task 2 has completed")
				return
			}

			for nIndex, n := range numbers {
				for mIndex, m := range numbers[nIndex+1:] {
					if n+m+r == 2020 {
						log.Printf(
							"[%d:%d, %d:%d, %d:%d] %d * %d * %d = %d",
							nIndex, numbers[nIndex],
							mIndex+nIndex+1, numbers[mIndex+1],
							len(numbers)+1, r,
							n, m, r, n*m*r,
						)
					}
				}
			}
			numbers = append(numbers, r)
		}
	}
}
