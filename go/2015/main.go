package main

import (
	"flag"
	"fmt"
	"os"
	"slices"

	"github.com/phortheman/AdventOfCode_2015_go/solutions/day01"
	"github.com/phortheman/AdventOfCode_2015_go/solutions/day02"
	"github.com/phortheman/AdventOfCode_2015_go/solutions/day03"
	"github.com/phortheman/AdventOfCode_2015_go/solutions/day04"
	"github.com/phortheman/AdventOfCode_2015_go/solutions/day05"
	"github.com/phortheman/AdventOfCode_2015_go/solutions/day06"
)

var (
	inputs = []string{
		"inputs/day_01.txt",
		"inputs/day_02.txt",
		"inputs/day_03.txt",
		"inputs/day_04.txt",
		"inputs/day_05.txt",
		"inputs/day_06.txt",
	}
)

var (
	specificDays = false
	days         = make([]bool, len(inputs))
)

func init() {
	// Dynamically add the flags as more solutions are provided
	for i := 0; i < len(days); i++ {
		flag.BoolVar(&days[i], fmt.Sprint(i+1), false, fmt.Sprintf("Run day %02d", i+1))
	}

	flag.Parse()

	// If any flags are provided then run only specific days
	specificDays = slices.Contains(days, true)
}

func main() {
	var day, part1, part2 int
	for i, input := range inputs {
		day++
		// If a specific day is specified and this is not one of those days, skip
		if specificDays && !days[i] {
			continue
		}

		content, err := os.ReadFile(input)
		if err != nil {
			fmt.Printf("\nMissing day %d input file. Make sure you are saving the puzzle input like this 'inputs/day_%02d.txt'\n", day, day)
			continue
		}

		switch day {
		case 1:
			part1, part2 = day01.Solver(content)

		case 2:
			part1, part2 = day02.Solver(string(content))

		case 3:
			part1, part2 = day03.Solver(content)

		case 4:
			part1, part2 = day04.Solver(string(content))

		case 5:
			part1, part2 = day05.Solver(string(content))

		case 6:
			part1, part2 = day06.Solver(string(content))
		}

		fmt.Printf("\nDay %d	Part 1: %d\n", day, part1)
		fmt.Printf("Day %d	Part 2: %d\n", day, part2)

		part1, part2 = 0, 0
	}
}
