package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"time"
)

const relative_input string = "../../../inputs/2015/05/input.txt"

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
	var part1, part2 int

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if isNicePart1(line) {
			part1 += 1
		}
		if isNicePart2(line) {
			part2 += 1
		}
	}

	return part1, part2
}

func isVowel[T rune | byte](char T) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}

func isNicePart1(input string) bool {
	var prevChar rune = rune(input[0])
	var vowelCount int = 0
	var doubleChars = false
	var blacklist = []string{"ab", "cd", "pq", "xy"}

	if isVowel(prevChar) {
		vowelCount += 1
	}

	for _, char := range input[1:] {
		if char == prevChar {
			doubleChars = true
		}

		if isVowel(char) {
			vowelCount += 1
		}

		combined := string([]rune{prevChar, char})
		if slices.Contains(blacklist, combined) {
			return false
		}
		prevChar = char
	}

	if doubleChars && vowelCount >= 3 {
		return true
	}

	return false
}

func isNicePart2(input string) bool {
	var prevprevChar rune = 0
	var prevChar rune = 0
	var combined string

	var pairCounts = make(map[string]int)
	var pairPosition = make(map[string]int)

	var doubleCharsWithSeparator = false

	for i, char := range input {
		if i == 0 {
			prevChar = rune(input[0])
			continue
		}

		combined = string([]rune{prevChar, char})

		if prevprevChar == char {
			doubleCharsWithSeparator = true
		}

		// If this pair has been found before get its idx. Otherwise it is the first pair
		if idx, ok := pairPosition[combined]; ok {
			// Because these are pairs only see if the cached idx is one less than current idx
			if idx+1 != i {
				pairCounts[combined] += 1
			}
		} else {
			pairPosition[combined] = i
			pairCounts[combined] += 1
		}

		prevprevChar = prevChar
		prevChar = char
	}

	if doubleCharsWithSeparator {
		for _, count := range pairCounts {
			if count >= 2 {
				return true
			}
		}
	}

	return false
}
