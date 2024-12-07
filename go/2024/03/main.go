package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"time"
)

const relative_input string = "../../../inputs/2024/03/input.txt"
const sampleInput1 = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
const sampleInput2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

var bTest1 = false
var bTest2 = false
var inputFileName string

func init() {
	flag.BoolVar(&bTest1, "t1", false, "Run using the part 1 sample input")
	flag.BoolVar(&bTest2, "t2", false, "Run using the part 2 sample input")
	flag.StringVar(&inputFileName, "i", "",
		"Path to the puzzle input. "+
			"Default to using the internal relative path. "+
			"Pass 'stdin' to use it instead")
}

func main() {
	flag.Parse()
	var content []byte
	var err error
	if !bTest1 && !bTest2 {
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
	} else if bTest1 {
		content = []byte(sampleInput1)
	} else if bTest2 {
		content = []byte(sampleInput2)
	}

	start := time.Now()
	part1, part2 := Solver(string(content))
	duration := time.Since(start)

	fmt.Printf("Time: %v\nPart 1: %d\nPart 2: %d\n", duration, part1, part2)
}

func Solver(input string) (int, int) {
	partOne, partTwo := 0, 0
	readMul := true
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|(don't\(\)|do\(\))`)

	matches := re.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		if match[0] == "do()" {
			readMul = true
		} else if match[0] == "don't()" {
			readMul = false
		} else {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			partOne += num1 * num2
			if readMul {
				partTwo += num1 * num2
			}
		}
	}

	return partOne, partTwo
}
