package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const relative_input string = "../../../inputs/2024/01/input.txt"
const sampleInput = `3   4
4   3
2   5
1   3
3   9
3   3`

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
	part1, part2 := Solver(string(content))
	duration := time.Since(start)

	fmt.Printf("Time: %v\nPart 1: %d\nPart 2: %d\n", duration, part1, part2)
}

func sortInsert(data []int, val int) []int {
	idx := sort.SearchInts(data, val)
	if idx == len(data) {
		return append(data, val)
	}

	// Make room
	data = append(data[:idx+1], data[idx:]...)

	data[idx] = val
	return data
}

func Solver(input string) (int, int) {
	leftLocations := make([]int, 0, len(input))
	rightLocations := make([]int, 0, len(input))
	frequency := make(map[int]int)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		locationIDs := strings.Split(line, "   ")

		left, _ := strconv.Atoi(locationIDs[0])
		right, _ := strconv.Atoi(locationIDs[1])
		frequency[right] += 1

		leftLocations = sortInsert(leftLocations, left)
		rightLocations = sortInsert(rightLocations, right)
	}

	distance := 0
	similarity := 0
	for i := 0; i < len(leftLocations); i++ {
		diff := leftLocations[i] - rightLocations[i]

		if diff < 0 {
			diff *= -1
		}

		distance += diff
		similarity += leftLocations[i] * frequency[leftLocations[i]]
	}

	return distance, similarity
}
