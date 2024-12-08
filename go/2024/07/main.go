package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const relative_input string = "../../../inputs/2024/07/input.txt"
const sampleInput = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

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

func Solver(input string) (int, int) {
	partOne, partTwo := 0, 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	var backtrack func(idx int, curExpr string)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ": ")
		target, _ := strconv.Atoi(tokens[0])
		tokens = strings.Fields(tokens[1])
		solutionsPartOne := false
		solutionsPartTwo := false
		backtrack = func(idx int, curExpr string) {
			if idx == len(tokens) {
				if solutionsPartOne && solutionsPartTwo {
					return
				}

				if eval(curExpr, target) {
					if !strings.Contains(curExpr, "||") {
						solutionsPartOne = true
					}
					solutionsPartTwo = true
				}
				return
			}

			backtrack(idx+1,
				fmt.Sprintf("%s + %s", curExpr, tokens[idx]))

			backtrack(idx+1,
				fmt.Sprintf("%s * %s", curExpr, tokens[idx]))

			backtrack(idx+1,
				fmt.Sprintf("%s || %s", curExpr, tokens[idx]))
		}

		if len(tokens) > 0 {
			backtrack(1, tokens[0])
		}

		if solutionsPartOne {
			partOne += target
		}

		if solutionsPartTwo {
			partTwo += target
		}

	}
	return partOne, partTwo
}

var concatTimes = make([]time.Duration, 0)

// Pass a string expression to be evaluated. Return early if left ever is greater than the target
func eval(expression string, target int) bool {
	tokens := strings.Fields(expression)
	left, _ := strconv.Atoi(tokens[0])

	// If there are less than 3 tokens then this is just a single number
	if len(tokens) < 3 {
		return left == target
	}

	for i := 1; i < len(tokens)-1; i += 2 {
		operator := tokens[i]
		right, _ := strconv.Atoi(tokens[i+1])

		switch operator {
		case "+":
			left += right
		case "*":
			left *= right
		case "||":
			left, _ = strconv.Atoi(strconv.Itoa(left) + tokens[i+1])
		}

		if left > target {
			return false
		}
	}

	return left == target
}
