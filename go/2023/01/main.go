package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const relative_input string = "../../../inputs/2023/01/input.txt"

const sampleInputPart1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

const sampleInputPart2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

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
	var content []byte
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

func Solver(input string) (int, int) {
	part1, part2 := 0, 0
	if !bTest {
		scanner := bufio.NewScanner(strings.NewReader(input))
		for scanner.Scan() {
			line := scanner.Text()
			part1 += PartOne([]byte(line))
			part2 += PartTwo([]byte(line))
		}
	} else {
		// Part 1 sample input
		scanner := bufio.NewScanner(strings.NewReader(sampleInputPart1))
		for scanner.Scan() {
			line := scanner.Text()
			part1 += PartOne([]byte(line))
		}

		// Part 2 sample input
		scanner = bufio.NewScanner(strings.NewReader(sampleInputPart2))
		for scanner.Scan() {
			line := scanner.Text()
			part2 += PartTwo([]byte(line))
		}
	}

	return part1, part2
}

func PartOne(input []byte) int {
	var f, s uint8
	l, r := 0, len(input)-1
	for {
		if f == 0 {
			if IsDigit(input[l]) {
				f = input[l] - '0'
			} else {
				l++
			}
		}
		if s == 0 {
			if IsDigit(input[r]) {
				s = input[r] - '0'
			} else {
				r--
			}
		}
		if f != 0 && s != 0 {
			break
		}
	}
	return int((f * 10) + s)
}

// Same as part one but adding logic to translate digits spelled out to uint8 values
func PartTwo(input []byte) int {
	var f, s uint8
	l, r := 0, len(input)-1
	for {
		if f == 0 {
			if IsDigit(input[l]) {
				f = input[l] - '0'
			} else if v := CheckForSpelledOutDigit(input[l:]); v != 0 {
				f = v
			} else {
				l++
			}
		}
		if s == 0 {
			if IsDigit(input[r]) {
				s = input[r] - '0'
			} else if v := CheckForSpelledOutDigit(input[r:]); v != 0 {
				s = v
			} else {
				r--
			}
		}
		if f != 0 && s != 0 {
			break
		}
	}
	return int((f * 10) + s)
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// Using switch statment to see if it is worth doing a prefix search or not
func CheckForSpelledOutDigit(input []byte) uint8 {
	switch input[0] {
	case 'o':
		if bytes.HasPrefix(input, []byte("one")) {
			return 1
		}
	case 't':
		if bytes.HasPrefix(input, []byte("two")) {
			return 2
		} else if bytes.HasPrefix(input, []byte("three")) {
			return 3
		}
	case 'f':
		if bytes.HasPrefix(input, []byte("four")) {
			return 4
		} else if bytes.HasPrefix(input, []byte("five")) {
			return 5
		}
	case 's':
		if bytes.HasPrefix(input, []byte("six")) {
			return 6
		} else if bytes.HasPrefix(input, []byte("seven")) {
			return 7
		}
	case 'e':
		if bytes.HasPrefix(input, []byte("eight")) {
			return 8
		}
	case 'n':
		if bytes.HasPrefix(input, []byte("nine")) {
			return 9
		}
	}
	return 0
}
