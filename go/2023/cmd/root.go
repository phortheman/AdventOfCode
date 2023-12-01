package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Application for the Advent of Code 2023",
	Long: `This is an application for solving the Advent of Code 2023 puzzles using Go along with the Cobra framework.
Each day is a sub command and each take an input file.`,
	Example: "aoc [command] --file input.txt",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("file", "f", "", "The file path for the puzzle input")
}
