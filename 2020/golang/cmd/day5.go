package cmd

import (
	"github.com/obukhov/advent-of-code/2020/golang/lib"
	"github.com/spf13/cobra"
	"log"
	"os"
	"sync"
)

func init() {
	rootCmd.AddCommand(day5cmd)
}

var day5cmd = &cobra.Command{
	Use:   "day5",
	Short: "Day 5 tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day5()
	},
}



func day5() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed getting working dir: %v", err)
	}

	task1input := make(chan int, 100)
	task2input := make(chan int, 100)

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go day5task1(task1input, wg)
	go day5task2(task2input, wg)

	go func() {
		lib.ReadFile(
			wd+"/input/day5.txt",
			func (line string) {
				seatId := 0
				for _, b := range line {
					seatId = seatId << 1
					if b == 'B' || b == 'R' {
						seatId++
					}
				}


				task1input <- seatId
				task2input <- seatId
			},
			wg,
		)

		close(task1input)
		close(task2input)
	}()

	wg.Wait()
}

func day5task1(input chan int, wg *sync.WaitGroup) {
	max := 0
	for seatId := range input {
		if seatId > max {
			max = seatId
		}
	}

	log.Printf("Max seatId: %d", max)
	wg.Done()
}

func day5task2(input chan int, wg *sync.WaitGroup) {
	stringMap := make([]bool, 1024, 1024)
	for seatId := range input {
		stringMap[seatId] = true
	}

	seatsStarted := false
	mySeat := 0
	for seatId, taken := range stringMap {
		if taken && !seatsStarted {
			seatsStarted = true
		}

		if !taken && seatsStarted {
			mySeat = seatId
			break
		}
	}
	log.Printf("My seat: %d", mySeat)
	wg.Done()
}
