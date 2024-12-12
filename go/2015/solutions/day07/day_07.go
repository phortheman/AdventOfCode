package day07

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

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
