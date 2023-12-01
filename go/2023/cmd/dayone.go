package cmd

import (
	file "aoc23/internal"
	"bytes"
	"fmt"
	"strconv"

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
	var total int
	var part2 int
	for _, line := range content {
		total += DayOnePartOne(line)
		part2 += DayOnePartTwo(line)
	}
	fmt.Println(total)
	fmt.Println(part2)
}

func DayOnePartOne(input []byte) int {
	l, r := 0, len(input)-1
	f, s := "", ""

	for f == "" || s == "" {
		if f == "" && IsDigit(input[l]) {
			f = string(input[l])
		}
		if s == "" && IsDigit(input[r]) {
			s = string(input[r])
		}
		l++
		r--
		if f != "" && s != "" {
			break
		}
		if r < 0 || l >= len(input) {
			break
		}
	}
	digit, _ := strconv.ParseInt(f+s, 10, 32)
	return int(digit)
}

func DayOnePartTwo(input []byte) int {
	f, s := 0, 0
	l, r := 0, len(input)-1
	for {
		if f == 0 {
			if IsDigit(input[l]) {
				f = int(input[l] - 48)
			} else if v := CheckForSpelledOutDigit(input[l:]); v != 0 {
				f = v
			} else {
				l++
			}
		}
		if s == 0 {
			if IsDigit(input[r]) {
				s = int(input[r] - 48)
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
	value := (f * 10) + s
	return value
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func CheckForSpelledOutDigit(input []byte) int {
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
