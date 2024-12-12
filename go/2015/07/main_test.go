package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`
	expected := map[string]uint16{
		"d": 72,
		"e": 507,
		"f": 492,
		"g": 114,
		"h": 65412,
		"i": 65079,
		"x": 123,
		"y": 456,
	}
	results := make(map[string]uint16)
	instr := make(map[string]string)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		instr[tokens[len(tokens)-1]] = strings.Join(tokens, " ")
	}

	// Just to make sure every wire is calculated
	_ = emulate(instr, results, "i")
	_ = emulate(instr, results, "h")
	_ = emulate(instr, results, "g")
	_ = emulate(instr, results, "f")
	_ = emulate(instr, results, "e")
	_ = emulate(instr, results, "d")

	for key, value := range results {
		if value != expected[key] {
			t.Errorf("Wire %s - Expected %d and got %d", key, expected[key], value)
		}
	}
	fmt.Println(results)
}
