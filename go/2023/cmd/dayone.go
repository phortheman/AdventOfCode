package cmd

import (
	file "aoc23/internal"
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
	for _, line := range content {
		total += DayOnePartOne(line)
	}
	fmt.Println(total)
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

func IsDigit(c byte) bool {
	return c >= 48 && c <= 57
}
