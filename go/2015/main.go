package main

import (
	"fmt"
	"os"

	"github.com/phortheman/AdventOfCode_2015_go/solutions"
)

var (
	inputs = []string{
		"inputs/day_01.txt",
		"inputs/day_02.txt",
		"inputs/day_03.txt",
	}
)

func main() {
	//TODO: parse args to run only specific day otherwise run every day
	for i, input := range inputs {
		day := i + 1
		var part1, part2 int
		content, err := os.ReadFile(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading file:", err)
			os.Exit(1)
		}

		switch day {
		case 1:
			part1, part2 = solutions.Day01Solver(content)

		case 2:
			part1, part2 = solutions.Day02Solver(string(content))

		case 3:
			part1, part2 = solutions.Day03Solver(content)
		}

		fmt.Printf("\nDay %d	Part 1: %d\n", day, part1)
		fmt.Printf("Day %d	Part 2: %d\n", day, part2)

	}

}
