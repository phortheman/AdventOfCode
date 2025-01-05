package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"time"
)

const relative_input string = "../../../inputs/2023/04/input.txt"
const sampleInput = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

var bTest = false
var inputFileName string

func init() {
	flag.BoolVar(&bTest, "t", false, "Run using the sample input")
	flag.StringVar(&inputFileName, "i", "",
		"Path to the puzzle input. "+
			"Default to using the internal relative path. "+
			"Pass 'stdin' to use it instead")
}

func main() {
	flag.Parse()
	var content []byte = []byte(sampleInput)
	var err error
	if !bTest {
		switch inputFileName {
		case "stdin":
			content, err = io.ReadAll(os.Stdin)
		case "":
			content, err = os.ReadFile(relative_input)
		default:
			content, err = os.ReadFile(inputFileName)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to read file: %v", err)
			os.Exit(66)
		}
	}

	start := time.Now()
	part1, part2 := Solver(content)
	duration := time.Since(start)

	fmt.Printf("Time: %v\nPart 1: %d\nPart 2: %d\n", duration, part1, part2)
}

func Solver(input []byte) (int, int) {
	newInput := bytes.Split(input, []byte("\n"))

	part1, part2 := 0, 0
	copies := make([]int, len(input))
	for i, line := range newInput {
		if len(line) == 0 {
			continue
		}

		_, t, _ := bytes.Cut(line, []byte(": "))
		winning, drawn, _ := bytes.Cut(t, []byte(" | "))
		var winningNums []int
		for _, c := range bytes.Split(winning, []byte(" ")) {
			// Account for double spaces
			if len(c) == 0 {
				continue
			}
			winningNums = append(winningNums, BytesToInt(c))
		}
		var points, count int = 0, 0
		copies[i]++
		for _, c := range bytes.Split(drawn, []byte(" ")) {
			// Account for double spaces
			if len(c) == 0 {
				continue
			}
			n := BytesToInt(c)
			if slices.Index(winningNums, n) != -1 {
				points *= 2
				if points == 0 {
					points = 1
				}
				count++
				if i+count >= len(copies) {
					break
				}
				copies[i+count] += copies[i]
			}
		}
		part1 += points
		part2 += copies[i]
	}
	return part1, part2
}

// Assumes bytes are numeric and works left to right
func BytesToInt(b []byte) int {
	var v int
	for _, c := range b {
		v = v*10 + int(c) - '0'
	}
	return v
}
