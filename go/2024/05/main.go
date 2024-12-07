package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

const relative_input string = "../../../inputs/2024/05/input.txt"
const sampleInput = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

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

	// NOTE: Default to converting puzzle input into a string representation of the whole file.
	// Change as needed for the puzzle
	start := time.Now()
	part1, part2 := Solver(string(content))
	duration := time.Since(start)

	fmt.Printf("Time: %v\nPart 1: %d\nPart 2: %d\n", duration, part1, part2)
}

func Solver(input string) (int, int) {
	partOne, partTwo := 0, 0

	re := regexp.MustCompile(`(\d+)\|(\d+)|.*,*`)
	matches := re.FindAllStringSubmatch(input, -1)

	pageOrder := make(map[string][]string)

	for _, match := range matches {
		if strings.Contains(match[0], "|") {
			pageOrder[match[1]] = append(pageOrder[match[1]], match[2])
		}
		if strings.Contains(match[0], ",") {
			pageUpdates := strings.Split(match[0], ",")
			if isValidOrder(pageOrder, pageUpdates) {
				partOne += getMiddle(pageUpdates)
			} else {
				orderUpdated := make([]string, 0, len(pageUpdates))
				orderUpdated = append(orderUpdated, pageUpdates...)
				slices.SortFunc(orderUpdated, func(a string, b string) int {
					before := pageOrder[a]
					if len(before) == 0 {
						return 0
					}

					if slices.Contains(before, b) {
						return -1
					}

					return 1
				})
				partTwo += getMiddle(orderUpdated)
			}
		}
	}

	return partOne, partTwo
}

func isValidOrder(before map[string][]string, updates []string) bool {
	for idx, update := range updates {
		// Check if any of the rules are actually after it
		for _, rule := range before[update] {
			if slices.Contains(updates[:idx], rule) {
				// This would be false because the number expected to be before was actually after
				return false
			}
		}
	}

	return true
}

func getMiddle(slice []string) int {
	middle, _ := strconv.Atoi(slice[len(slice)/2])
	return middle
}
