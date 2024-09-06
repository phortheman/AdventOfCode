package main

import (
	"fmt"
	"os"

	"github.com/phortheman/AdventOfCode_2015_go/solutions"
)

var (
	input = [25]string{
		"inputs/day_01.txt",
	}
)

func main() {
	//TODO: parse args to run only specific day otherwise run every day
	content, err := os.ReadFile("inputs/day_01.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading file:", err)
		os.Exit(1)
	}
	part1, part2, err := solutions.Day01Solver(string(content))
	if err != nil {
		fmt.Fprintln(os.Stderr, "solver error:", err)
		os.Exit(1)
	}

	fmt.Println("Day 1 Part 1:", part1)
	fmt.Println("Day 1 Part 2:", part2)

}
