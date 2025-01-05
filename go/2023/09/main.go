package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const relative_input string = "../../../inputs/2023/09/input.txt"
const sampleInput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

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
	nextTotal, prevTotal := 0, 0
	for _, line := range newInput {
		if len(line) == 0 {
			continue
		}
		v := SplitIntoInts(line)
		next, prev := Extrapolate(v)
		nextTotal += next
		prevTotal += prev
	}
	return nextTotal, prevTotal
}

func SplitIntoInts(input []byte) []int {
	splitedBytes := bytes.Split(input, []byte(" "))
	output := make([]int, 0, len(splitedBytes))
	for _, b := range splitedBytes {
		output = append(output, BytesToInt(b))
	}
	return output
}

func BytesToInt(input []byte) int {
	s := string(input)
	v, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Non-numeric value:", s)
		os.Exit(1)
	}
	return v
}

// Returns the extrapolated previous and next value in a sequence of ints
func Extrapolate(input []int) (int, int) {
	diffs := make([]int, 0, len(input)-1)
	bAllZeros := true
	for i := 0; i < len(input)-1; i++ {
		diff := input[i+1] - input[i]
		if bAllZeros && diff != 0 {
			bAllZeros = false
		}
		diffs = append(diffs, diff)
	}
	var nextDiff, prevDiff = 0, 0
	if !bAllZeros {
		nextDiff, prevDiff = Extrapolate(diffs)
	}
	return input[len(input)-1] + nextDiff, input[0] - prevDiff
}
