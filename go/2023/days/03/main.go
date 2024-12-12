package main

import (
	file "aoc23/internal"
	"bytes"
	"fmt"
	"os"
)

var SYMBOLS string = "!@#$%^&*()-_=+<>?/|]}[{;:'"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "<file_path>")
		os.Exit(1)
	}

	content, err := file.Read_File_Into_Memory(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file into memory: ", err)
	}

	var part1Total uint = Part1Solver(content)
	var part2Total uint = Part2Solver(content)

	fmt.Println(part1Total)
	fmt.Println(part2Total)
}

func Part1Solver(input [][]byte) uint {
	var total uint = 0
	for i, line := range input {
		var curVal uint = 0
		startPos := -1
		for j := 0; j <= len(line); j++ {
			if j == len(line) {
				if startPos != -1 {
					if EvalPartNumber(startPos, j-1, i, input) {
						total += curVal
					}
					curVal = 0
					startPos = -1
				}
				continue
			}
			if IsDigit(line[j]) {
				if startPos == -1 {
					startPos = j
				}
				curVal = BuildUintFromByte(line[j], curVal)
				continue
			}
			if curVal != 0 {
				if EvalPartNumber(startPos, j-1, i, input) {
					total += curVal
				}
				curVal = 0
				startPos = -1
			}
		}
	}
	return total
}

func Part2Solver(input [][]byte) uint {
	var total uint = 0
	for i, line := range input {
		for j := 0; j < len(line); j++ {
			if line[j] == '*' {
				total += EvalGear(j, i, input)
			}
		}
	}
	return total

}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func BuildUintFromByte(c byte, v uint) uint {
	return uint(v*10 + uint(c) - '0')
}

// Assumes no alpha characters or other whitespace/non-printable characters
func IsSymbol(c byte) bool {
	return !IsDigit(c) && c != '.'
}

func EvalPartNumber(start, end, y int, input [][]byte) bool {
	xMin, xMax := start, end
	if xMin != 0 {
		xMin -= 1
		if IsSymbol(input[y][xMin]) {
			return true
		}
	}
	if xMax != len(input[y])-1 {
		xMax += 1
		if IsSymbol(input[y][xMax]) {
			return true
		}
	}
	if y != 0 {
		if bytes.ContainsAny(input[y-1][xMin:xMax+1], SYMBOLS) {
			return true
		}
	}
	if y != len(input)-1 {
		if bytes.ContainsAny(input[y+1][xMin:xMax+1], SYMBOLS) {
			return true
		}
	}
	return false
}

func EvalGear(x, y int, input [][]byte) uint {
	adjNums := []uint{}
	nums := []Ratio{}
	if y != 0 {
		nums = append(nums, CalculateNumbers(input[y-1])...)
	}
	nums = append(nums, CalculateNumbers(input[y])...)
	if y != len(input)-1 {
		nums = append(nums, CalculateNumbers(input[y+1])...)
	}
	for _, num := range nums {
		if num.IsAdj(x-1) || num.IsAdj(x) || num.IsAdj(x+1) {
			adjNums = append(adjNums, num.value)
		}
	}
	if len(adjNums) == 2 {
		return adjNums[0] * adjNums[1]
	}
	return 0
}

type Ratio struct {
	start int
	end   int
	value uint
}

func (r *Ratio) IsAdj(n int) bool {
	return n >= r.start && n <= r.end
}

func CalculateNumbers(input []byte) []Ratio {
	output := []Ratio{}
	for i := 0; i < len(input); i++ {
		if IsDigit(input[i]) {
			curRatio := Ratio{}
			curRatio.start = i
			for IsDigit(input[i]) {
				curRatio.value = BuildUintFromByte(input[i], curRatio.value)
				i++
				if i == len(input) {
					break
				}
			}
			curRatio.end = i - 1
			i--
			output = append(output, curRatio)
		}
	}
	return output
}
