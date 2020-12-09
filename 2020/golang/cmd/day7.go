package cmd

import (
	"github.com/obukhov/advent-of-code/2020/golang/lib"
	"github.com/spf13/cobra"
	"log"
	"os"
	"regexp"
	"strconv"
	"sync"
)

func init() {
	rootCmd.AddCommand(day7cmd)
}

var day7cmd = &cobra.Command{
	Use:   "day7",
	Short: "Day 7 tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day7()
	},
}

type bag string
type set struct {
	num int
	bag bag
}
type rule struct {
	outer  bag
	contained []set
}

func day7() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed getting working dir: %v", err)
	}

	expFull, err := regexp.Compile("^(.+?) bags contain (no other bags|\\d+ .+? bags?(?:, \\d+ .+? bags?)*)\\.$")
	if err != nil {
		log.Fatal(err)
	}
	expInner, err := regexp.Compile("(\\d+) (.+?) bags?")
	if err != nil {
		log.Fatal(err)
	}

	task1input := make(chan rule, 100)
	task2input := make(chan rule, 100)

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go day7task1(task1input, wg)
	go day7task2(task2input, wg)

	go func() {
		lib.ReadFileWg(
			wd+"/input/day7.txt",
			func (line string) {
				subMatches := expFull.FindStringSubmatch(line)
				if len(subMatches) != 3 {
					log.Fatalf("Error parsing '%s'", line)
				}

				rule := rule{outer: bag(subMatches[1]), contained: make([]set, 0)}
				if subMatches[1] != "no other bags" {
					innerMatches := expInner.FindAllStringSubmatch(subMatches[2], 10)
					for _, m := range innerMatches{
						i, err := strconv.Atoi(m[1])
						if err != nil {
							log.Fatalf("Error parsing '%s' to number", m[1])
						}
						rule.contained = append(rule.contained, set{i, bag(m[2])})
					}
				}

				task1input <- rule
				task2input <- rule
			},
			wg,
		)

		close(task1input)
		close(task2input)
	}()

	wg.Wait()
}

func day7task1(input chan rule, wg *sync.WaitGroup) {
	var reverseMap = make(map[bag][]bag)

	for r := range input {
		for _, c := range r.contained {
			if rules, found := reverseMap[c.bag]; found {
				rules = append(rules, r.outer)
				reverseMap[c.bag] = rules
			} else {
				rules = []bag{r.outer}
				reverseMap[c.bag] = rules
			}
		}
	}

	//for k ,val := range reverseMap {
	//	log.Printf("%s %v", k, val)
	//}

	parentBags := findParents(reverseMap, bag("shiny gold"))

	//log.Println(parentBags)
	log.Println("Count outer bags: ", len(parentBags))
	wg.Done()
}

func findParents(reverseMap map[bag][]bag, node bag) map[bag]bool {
	result := make(map[bag]bool, 0)
	if outerBags, found := reverseMap[node]; found {
		for _,	cur := range outerBags {
			result[cur] = true
			for k, v := range findParents(reverseMap, cur){
				result[k] = v
			}
		}
	}

	return result
}

func day7task2(input chan rule, wg *sync.WaitGroup) {
	var descMap = make(map[bag][]set)

	for r := range input {
		descMap[r.outer] = r.contained
	}

	for k ,val := range descMap {
		log.Printf("%s %v", k, val)
	}

	count := countSubtree(descMap, bag("shiny gold"))
	log.Println("Count inner bags: ", count)

	wg.Done()
}

func countSubtree(descMap map[bag][]set, node bag) int {
	count := 0
	if subBags, found :=  descMap[node]; found {
		for _, subBag := range subBags {
			count += subBag.num * (1 + countSubtree(descMap, subBag.bag))
		}
	}

	return count
}