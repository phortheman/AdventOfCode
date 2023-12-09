package main

import (
	file "aoc23/internal"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var EXAMPLE string = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func main() {
	var content [][]byte
	if len(os.Args) != 2 {
		content = file.Read_String_Into_Byte_Slice(EXAMPLE)
	} else {
		var err error
		content, err = file.Read_File_Into_Memory(os.Args[1])
		if err != nil {
			fmt.Println("Error reading file into memory: ", err)
			os.Exit(1)
		}
	}

	part1Total, part2Total := Solver(content)

	fmt.Println(part1Total)
	fmt.Println(part2Total)
}

func Solver(input [][]byte) (int, int) {
	nextTotal, prevTotal := 0, 0
	for _, line := range input {
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
