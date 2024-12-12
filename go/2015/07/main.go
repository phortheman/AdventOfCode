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

const relative_input string = "../../../inputs/2015/07/input.txt"

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

	instr := make(map[string]string)
	signals := make(map[string]uint16)
	scanner := bufio.NewScanner(strings.NewReader(input))

	// Parse the lines and setup the map to store the instructions for the wire signal
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		instr[tokens[len(tokens)-1]] = strings.Join(tokens, " ")
	}

	part1 = int(emulate(instr, signals, "a"))
	instr["b"] = fmt.Sprintf("%d -> b", part1)

	for k := range signals {
		delete(signals, k)
	}

	part2 = int(emulate(instr, signals, "a"))

	return part1, part2
}

func toUint16(str string) uint16 {
	value, _ := strconv.Atoi(str)
	return uint16(value)
}

func emulate(instr map[string]string, signals map[string]uint16, target string) uint16 {
	signal, ok := signals[target]
	if ok {
		return signal
	}

	tokens := strings.Split(instr[target], " ")

	if len(tokens) <= 1 {
		return toUint16(target)
	}

	switch tokens[1] {
	case "->":
		a, ok := signals[tokens[0]]
		if !ok {
			a = emulate(instr, signals, tokens[0])
		}
		signals[target] = a
		return signals[target]

	case "AND":
		a, ok := signals[tokens[0]]
		if !ok {
			a = emulate(instr, signals, tokens[0])
		}

		b, ok := signals[tokens[2]]
		if !ok {
			b = emulate(instr, signals, tokens[2])
		}

		signals[target] = a & b
		return signals[target]

	case "OR":
		a, ok := signals[tokens[0]]
		if !ok {
			a = emulate(instr, signals, tokens[0])
		}

		b, ok := signals[tokens[2]]
		if !ok {
			b = emulate(instr, signals, tokens[2])
		}
		signals[target] = a | b

		return signals[target]

	case "LSHIFT":
		a, ok := signals[tokens[0]]
		if !ok {
			a = emulate(instr, signals, tokens[0])
		}

		shift := toUint16(tokens[2])
		signals[target] = a << shift

		return signals[target]

	case "RSHIFT":
		a, ok := signals[tokens[0]]
		if !ok {
			a = emulate(instr, signals, tokens[0])
		}

		shift := toUint16(tokens[2])

		signals[target] = a >> shift
		return signals[target]

	default: // NOT
		a, ok := signals[tokens[1]]
		if !ok {
			a = emulate(instr, signals, tokens[1])
		}
		signals[target] = ^a
		return signals[target]
	}
}
