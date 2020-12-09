package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "advent2020",
	Short: "Advent of code tasks, add subcommand taskN",
	Long: "See for more info: https://adventofcode.com/",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Advent of code 2020")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}