package cmd

import (
	"github.com/obukhov/advent-of-code/2020/golang/lib"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"log"
)

func init() {
	rootCmd.AddCommand(day9cmd)
}


var day9cmd = &cobra.Command{
	Use:   "day9",
	Short: "Day 9 tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		preambleLen := 25

		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed getting working dir: %v", err)
		}

		numbers := make(chan int)
		go lib.ReadFile(
			wd+"/input/day9.txt",
			func (line string) {
				num, err := strconv.Atoi(line)
				if err != nil {
					log.Fatal(err)
				}

				numbers <- num
			},
			func () {
				close(numbers)
			},
		)

		preamble := make([]int, 0, preambleLen)
		set := make([]int, 0, 0)

		for i := 0; i < preambleLen; i++ {
			num := <- numbers
			preamble = append(preamble, num)
			set = append(set, num)
		}

		invalidNum := 0
		for num := range numbers {
			if findSum(preamble, num) {
				preamble = append(preamble[1:], num)
				set = append(set, num)
			} else {
				log.Printf("Invalid number: %d", num)
				invalidNum = num
				break
			}
		}

		sequence := make([]int, 0)
		sum := 0
		for _, v := range set {
			sequence = append(sequence, v)
			sum += v

			for sum > invalidNum {
				sum -= sequence[0]
				sequence = sequence[1:]
			}

			if sum == invalidNum {
				break
			}
		}

		log.Println("Sequence of numbers: ", sequence)
		min, max := sequence[0], sequence[0]
		for _, n := range sequence[1:] {
			if n < min {
				min = n
			} else if n > max {
				max = n
			}
		}

		log.Printf("%d + %d = %d", min, max, min+ max)

	},
}

func findSum(nums []int, sum int ) bool {
	for i, n := range nums {
		for _, m := range nums[i+1:] {
			if n+m == sum {
				return true
			}
		}
	}

	return false
}