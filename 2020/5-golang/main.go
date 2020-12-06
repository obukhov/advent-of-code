package main

import (
	"github.com/obukhov/advent-of-code/common"
	"log"
	"os"
	"sync"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed getting working dir: %v", err)
	}

	task1input := make(chan int, 100)
	task2input := make(chan int, 100)

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go task1(task1input, wg)
	go task2(task2input, wg)

	go func() {
		common.ReadFile(
			wd + "/input.txt",
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

func task1(input chan int, wg *sync.WaitGroup) {
	max := 0
	for seatId := range input {
		if seatId > max {
			max = seatId
		}
	}

	log.Printf("Max seatId: %d", max)
	wg.Done()
}

func task2(input chan int, wg *sync.WaitGroup) {
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
