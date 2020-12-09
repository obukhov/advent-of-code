package cmd

import (
	"github.com/obukhov/advent-of-code/2020/golang/lib"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
)

func init() {
	rootCmd.AddCommand(day8cmd)
}

const (
	NOP = "nop"
	JMP = "jmp"
	ACC = "acc"
)

type instruction struct {
	command string
	argument int
}

var day8cmd = &cobra.Command{
	Use:   "day8",
	Short: "Day 8 tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed getting working dir: %v", err)
		}

		instructions := make(chan instruction)
		go lib.ReadFile(
			wd+"/input/day8.txt",
			func (line string) {
				command := strings.Split(line, " ")
				arg, err := strconv.Atoi(command[1])
				if err != nil {
					log.Fatal(err)
				}

				instructions <- instruction{command[0], arg}
			},
			func () {
				close(instructions)
			},
		)

		program := make([]instruction, 0)
		for instruction := range instructions {
			program = append(program, instruction)
		}

		log.Printf("Loaded program with %d instructions", len(program))
		acc, terminated := run(program)
		log.Printf("Accumulated %d, terminated: %v", acc, terminated)

		for k, _ := range program {
			newProgram, isPatched := patch(program, k)
			if !isPatched {
				continue
			}

			acc, terminated := run(newProgram)
			if terminated {
				log.Printf("Program is terminated when patched in pos %d, accumulator %d", k, acc)
			}
		}
	},
}

func patch(program []instruction, pos int) (newProgram []instruction, isPatched bool) {
	newProgram = append([]instruction{}, program...)

	switch newProgram[pos].command {
	case NOP:
		newProgram[pos].command = JMP
		isPatched = true
	case JMP:
		newProgram[pos].command = NOP
		isPatched = true
	}

	return
}
func run(program []instruction) (acc int, terminated bool) {
	var (
		programLen = len(program)
		visited = make([]bool, programLen, programLen)
		ptr = 0
	)

	for {
		cur := program[ptr]
		switch cur.command {
		case NOP:
			ptr++
		case ACC:
			acc = acc + cur.argument
			ptr++
		case JMP:
			ptr = ptr + cur.argument
		}

		if ptr >= programLen {
			terminated = true
			return
		}

		if visited[ptr] {
			return
		} else {
			visited[ptr] = true
		}
	}
}