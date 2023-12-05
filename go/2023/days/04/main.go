package main

import (
	file "aoc23/internal"
	"bytes"
	"fmt"
	"os"
	"slices"
)

var EXAMPLE string = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func main() {
	var content [][]byte
	if len(os.Args) != 2 {
		content = file.Read_String_Into_Byte_Slice(EXAMPLE)
	} else {
		var err error
		content, err = file.Read_File_Into_Memory(os.Args[1])
		if err != nil {
			fmt.Println("Error reading file into memory: ", err)
			os.Exit(1)
		}
	}

	var part1Total, part2Total uint = PartSolver(content)

	fmt.Println(part1Total)
	fmt.Println(part2Total)
}

func PartSolver(input [][]byte) (uint, uint) {
	var part1Total uint = 0
	var part2Total uint = 0
	cards := make([]Card, len(input))
	for i, line := range input {
		t := bytes.TrimPrefix(line, []byte("Card "))
		_, t, _ = bytes.Cut(t, []byte(": "))
		winning, drawn, _ := bytes.Cut(t, []byte(" | "))
		var winningNums []uint
		for _, c := range bytes.Split(winning, []byte(" ")) {
			if len(c) == 0 {
				continue
			}
			winningNums = append(winningNums, BytesToUint(c))
		}
		count := 0
		cards[i].copies++
		for _, c := range bytes.Split(drawn, []byte(" ")) {
			if len(c) == 0 {
				continue
			}
			n := BytesToUint(c)
			if slices.Index(winningNums, n) != -1 {
				cards[i].points *= 2
				if cards[i].points == 0 {
					cards[i].points = 1
				}
				count++
				if i+count >= len(cards) {
					break
				}
				cards[i+count].copies += cards[i].copies
			}
		}
		part1Total += cards[i].points
		part2Total += cards[i].copies
	}
	return part1Total, part2Total
}

// Assumes byte is numeric and works left to right
func BytesToUint(b []byte) uint {
	var v uint
	for _, c := range b {
		v = uint(v*10 + uint(c) - '0')
	}
	return v
}

type Card struct {
	points uint
	copies uint
}
