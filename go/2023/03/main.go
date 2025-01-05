package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

const relative_input string = "../../../inputs/2023/03/input.txt"
const sampleInput = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

var bTest = false
var inputFileName string

var SYMBOLS string = "!@#$%^&*()-_=+<>?/|]}[{;:'"

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
	// If reading from a file a new line feed is inserted at the end of the byte slice. See if we need to adjust for that
	adjustment := 0
	if len(input) > 0 && input[len(input)-1] != byte('\n') {
		adjustment = 1
	}

	// Count how many lines are in the puzzle input
	lineCount := bytes.Count(input, []byte("\n")) + adjustment

	// Create the 2D slice of bytes to represent the grid
	newInput := make([][]byte, lineCount)
	fmt.Println("Line count:", lineCount)

	// Scan the puzzle input line by line
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for i := 0; scanner.Scan(); i++ {
		// Allocate the exact line length before copying the bytes over.
		// The scanner.Bytes() has a capacity of 4096 so this will reduce the memory usage for shorter lines
		size := len(scanner.Bytes())
		if size > 0 {
			newInput[i] = make([]byte, len(scanner.Bytes()))
			copy(newInput[i], scanner.Bytes())
		}
	}

	// Feed the new 2D puzzle input into each part solver
	part1 := Part1(newInput)
	part2 := Part2(newInput)

	return part1, part2
}

func Part1(input [][]byte) int {
	var total int = 0
	for i, line := range input {
		var curVal int = 0
		startPos := -1
		for j := 0; j <= len(line); j++ {
			if j == len(line) {
				if startPos != -1 {
					if EvalPartNumber(startPos, j-1, i, input) {
						total += curVal
					}
					curVal = 0
					startPos = -1
				}
				continue
			}
			if IsDigit(line[j]) {
				if startPos == -1 {
					startPos = j
				}
				curVal = BuildIntFromByte(line[j], curVal)
				continue
			}
			if curVal != 0 {
				if EvalPartNumber(startPos, j-1, i, input) {
					total += curVal
				}
				curVal = 0
				startPos = -1
			}
		}
	}
	return total
}

func Part2(input [][]byte) int {
	var total int = 0
	for i, line := range input {
		for j := 0; j < len(line); j++ {
			if line[j] == '*' {
				total += EvalGear(j, i, input)
			}
		}
	}
	return total

}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func BuildIntFromByte(c byte, v int) int {
	return int(v*10 + int(c) - '0')
}

// Assumes no alpha characters or other whitespace/non-printable characters
func IsSymbol(c byte) bool {
	return !IsDigit(c) && c != '.'
}

func EvalPartNumber(start, end, y int, input [][]byte) bool {
	xMin, xMax := start, end
	if xMin != 0 {
		xMin -= 1
		if IsSymbol(input[y][xMin]) {
			return true
		}
	}
	if xMax != len(input[y])-1 {
		xMax += 1
		if IsSymbol(input[y][xMax]) {
			return true
		}
	}
	if y != 0 {
		if bytes.ContainsAny(input[y-1][xMin:xMax+1], SYMBOLS) {
			return true
		}
	}
	if y != len(input)-1 {
		if bytes.ContainsAny(input[y+1][xMin:xMax+1], SYMBOLS) {
			return true
		}
	}
	return false
}

func EvalGear(x, y int, input [][]byte) int {
	adjNums := []int{}
	nums := []Ratio{}
	if y != 0 {
		nums = append(nums, CalculateNumbers(input[y-1])...)
	}
	nums = append(nums, CalculateNumbers(input[y])...)
	if y != len(input)-1 {
		nums = append(nums, CalculateNumbers(input[y+1])...)
	}
	for _, num := range nums {
		if num.IsAdj(x-1) || num.IsAdj(x) || num.IsAdj(x+1) {
			adjNums = append(adjNums, num.value)
		}
	}
	if len(adjNums) == 2 {
		return adjNums[0] * adjNums[1]
	}
	return 0
}

type Ratio struct {
	start int
	end   int
	value int
}

func (r *Ratio) IsAdj(n int) bool {
	return n >= r.start && n <= r.end
}

func CalculateNumbers(input []byte) []Ratio {
	output := []Ratio{}
	for i := 0; i < len(input); i++ {
		if IsDigit(input[i]) {
			curRatio := Ratio{}
			curRatio.start = i
			for IsDigit(input[i]) {
				curRatio.value = BuildIntFromByte(input[i], curRatio.value)
				i++
				if i == len(input) {
					break
				}
			}
			curRatio.end = i - 1
			i--
			output = append(output, curRatio)
		}
	}
	return output
}
