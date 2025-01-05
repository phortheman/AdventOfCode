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

const relative_input string = "../../../inputs/2023/02/input.txt"
const sampleInput = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

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

// NOTE: Base structure for the solver. Should handle the line by line operation for those puzzles
func Solver(input string) (int, int) {
	var part1Total int
	var part2Total int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		part1Total += PartOne(string(line))
		part2Total += PartTwo(string(line))
	}
	return int(part1Total), int(part2Total)
}

func PartOne(s string) int {
	id, rounds := GetGameData(s)
	for _, round := range rounds {
		cubes := strings.Split(round, ", ")
		for _, cube := range cubes {
			count, color := SplitCubeData(cube)
			if color == "red" && count > 12 {
				return 0
			} else if color == "green" && count > 13 {
				return 0
			} else if color == "blue" && count > 14 {
				return 0
			}
		}
	}
	return id
}

func PartTwo(s string) int {
	var red, green, blue int
	_, rounds := GetGameData(s)
	for _, round := range rounds {
		cubes := strings.Split(round, ", ")
		for _, cube := range cubes {
			count, color := SplitCubeData(cube)
			if color == "red" && count > red {
				red = count
			} else if color == "green" && count > green {
				green = count
			} else if color == "blue" && count > blue {
				blue = count
			}
		}
	}
	return red * green * blue
}

func GetGameData(s string) (int, []string) {
	var p int8
	for {
		if IsDigit(s[p]) {
			break
		}
		p++
	}
	var id int
	for {
		if s[p] == ':' {
			p += 2
			break
		}
		id = id*10 + int(s[p]) - '0'
		p++
	}
	return id, strings.Split(s[p:], "; ")
}

func SplitCubeData(s string) (int, string) {
	d := strings.Split(s, " ")
	c, _ := strconv.Atoi(d[0])
	return c, d[1]
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
