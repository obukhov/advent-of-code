package cmd

import (
	"github.com/obukhov/advent-of-code/2020/golang/lib"
	"github.com/spf13/cobra"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func init() {
	rootCmd.AddCommand(day4cmd)
}

var day4cmd = &cobra.Command{
	Use:   "day4",
	Short: "Day 4 tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day4()
	},
}

func day4() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed getting working dir: %v", err)
	}

	task1input := make(chan map[string]string, 100)
	task2input := make(chan map[string]string, 100)

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go day4task1(task1input, wg)
	go day4task2(task2input, wg)

	go func() {
		lib.ReadFileWg(
			wd+"/input/day4.txt",
			func (line string) {
				var params = make(map[string]string)
				if line != "" {
					pairs := strings.Split(line, " ")
					for _, pair := range pairs {
						keyValue := strings.SplitN(pair, ":", 2)
						params[keyValue[0]] = keyValue[1]
					}
				}

				task1input <- params
				task2input <- params
			},
			wg,
		)

		close(task1input)
		close(task2input)
	}()

	wg.Wait()
}

func day4task1(input chan map[string]string, wg *sync.WaitGroup) {
	required := map[string]string{}
	validCount := 0
	required = map[string]string{
		"byr": "",
		"iyr": "",
		"eyr": "",
		"hgt": "",
		"hcl": "",
		"ecl": "",
		"pid": "",
	}
	missing := make(map[string]string)
	for k,v := range required {
		missing[k] = v
	}
	for pairs := range input {
		if len(pairs) == 0 {
			if len(missing) == 0 {
				validCount++
			}
			for k,v := range required {
				missing[k] = v
			}
		} else {
			for key,_ := range pairs {
				delete(missing, key)
			}
		}
	}

	if len(missing) == 0 {
		validCount++
	}

	log.Printf("Total valid passports: %d", validCount)
	wg.Done()
}

type validatorFunc func(in string) bool

func createYearValidator(min, max int) validatorFunc {
	return func(in string) bool {
		v, err := strconv.Atoi(in)
		if err != nil {
			return false
		}

		return min <= v && v <= max
	}
}

func createRegexpValidator(expString string) validatorFunc {
	exp, err := regexp.Compile(expString)
	if err != nil {
		log.Fatal(err)
	}

	return func(in string) bool {
		return exp.MatchString(in)
	}
}


func day4task2(input chan map[string]string, wg *sync.WaitGroup) {
	var (
		cmValidator = createYearValidator(150, 193)
		inValidator = createYearValidator(59, 76)
		task2validator = map[string]validatorFunc {
			"byr": createYearValidator( 1920, 2002),
			"iyr": createYearValidator( 2010, 2020),
			"eyr": createYearValidator( 2020, 2030),
			"hgt": func(in string) bool {
				switch {
				case strings.HasSuffix(in, "cm"):
					return cmValidator(in[0:len(in)-2])
				case strings.HasSuffix(in, "in"):
					return inValidator(in[0:len(in)-2])
				default:
					return false
				}
			},
			"hcl": createRegexpValidator("^#[0-9a-f]{6}$"),
			"ecl": createRegexpValidator("^(amb|blu|brn|gry|grn|hzl|oth)$"),
			"pid": createRegexpValidator("^\\d{9}$"),
		}
		required = map[string]string{
			"byr": "",
			"iyr": "",
			"eyr": "",
			"hgt": "",
			"hcl": "",
			"ecl": "",
			"pid": "",
		}
		missingOrInvalid = make(map[string]string)
		validCount       = 0
	)

	for k,v := range required {
		missingOrInvalid[k] = v
	}

	for pairs := range input {
		if len(pairs) == 0 {
			if len(missingOrInvalid) == 0 {
				validCount++
			}
			for k,v := range required {
				missingOrInvalid[k] = v
			}
		} else {
			for key, val := range pairs {
				if validator, ok := task2validator[key]; ok {
					if validator(val) {
						delete(missingOrInvalid, key)
					}
				} else {
					delete(missingOrInvalid, key)
				}
			}
		}
	}

	if len(missingOrInvalid) == 0 {
		validCount++
	}

	log.Printf("Total valid passports: %d", validCount)
	wg.Done()
}
