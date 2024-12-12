package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const relative_input string = "../../../inputs/2015/08/input.txt"

var bTest = false
var inputFileName string

func init() {
	flag.StringVar(&inputFileName, "i", "",
		"Path to the puzzle input. "+
			"Default to using the internal relative path. "+
			"Pass 'stdin' to use it instead")
}

func main() {
	flag.Parse()
	var content []byte
	var err error
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

	start := time.Now()
	part1, part2 := Solver(string(content))
	duration := time.Since(start)

	fmt.Printf("Time: %v\nPart 1: %d\nPart 2: %d\n", duration, part1, part2)
}

func Solver(input string) (int, int) {
	part1 := 0
	part2 := 0

	chars, memory, encodedMemory := 0, 0, 0

	scan := bufio.NewScanner(strings.NewReader(input))
	for scan.Scan() {
		str := scan.Text()
		encodedStr := encodeString(str)
		chars += countChars(str)
		memory += len(str)
		encodedMemory += len(encodedStr)
	}

	part1 = memory - chars
	part2 = encodedMemory - memory

	return part1, part2
}

func countChars(input string) int {
	// Start with the number of characters minue the quote wrap
	count := len(input) - 2
	escaped := false

	for _, char := range input {
		// Remove the escape character
		if !escaped && char == '\\' {
			escaped = true
			count--
			continue
		}

		// This is a hex value so subtract the next 2 character is takes to define one
		if escaped && char == 'x' {
			count -= 2
		}

		escaped = false
	}
	return count
}

func encodeString(input string) string {
	var encodedString []rune = make([]rune, 0)
	for _, char := range input {
		if char == '"' || char == '\\' {
			encodedString = append(encodedString, '\\')
		}
		encodedString = append(encodedString, char)
	}
	return fmt.Sprintf(`"%s"`, string(encodedString))
}
