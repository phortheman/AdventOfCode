package cmd

import (
	file "aoc23/internal"
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
)

// dayoneCmd represents the dayone command
var dayoneCmd = &cobra.Command{
	Use:   "dayone",
	Short: "Solver for Day One",
	Run:   DayOneSolver,
}

func init() {
	rootCmd.AddCommand(dayoneCmd)
}

func DayOneSolver(cmd *cobra.Command, args []string) {
	filepath, _ := rootCmd.Flags().GetString("file")
	content, err := file.Read_File_Into_Memory(filepath)
	if err != nil {
		fmt.Println("Error reading file into memory: ", err)
	}
	var totalPart1 uint32
	var totalPart2 uint32
	for _, line := range content {
		totalPart1 += DayOnePartOne(line)
		totalPart2 += DayOnePartTwo(line)
	}
	fmt.Println("Part 1: ", totalPart1)
	fmt.Println("Part 2: ", totalPart2)
}

// Double pointer pinching in until both left and right have a value
func DayOnePartOne(input []byte) uint32 {
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
	return uint32((f * 10) + s)
}

// Same as part one but adding logic to translate digits spelled out to uint8 values
func DayOnePartTwo(input []byte) uint32 {
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
	return uint32((f * 10) + s)
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
