package main

import (
	file "aoc23/internal"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "<file_path>")
		os.Exit(1)
	}

	content, err := file.Read_File_Into_Memory(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file into memory: ", err)
	}

	var part1Total uint
	var part2Total uint
	for _, line := range content {
		part1Total += DayTwoPartOne(string(line))
		part2Total += DayTwoPartTwo(string(line))
	}
	fmt.Println("Part 1: ", part1Total)
	fmt.Println("Part 2: ", part2Total)
}

func DayTwoPartOne(s string) uint {
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

func DayTwoPartTwo(s string) uint {
	var red, green, blue uint
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

func GetGameData(s string) (uint, []string) {
	var p uint8
	for {
		if IsDigit(s[p]) {
			break
		}
		p++
	}
	var id uint
	for {
		if s[p] == ':' {
			p += 2
			break
		}
		id = uint(id*10 + uint(s[p]) - '0')
		p++
	}
	return id, strings.Split(s[p:], "; ")
}

func SplitCubeData(s string) (uint, string) {
	d := strings.Split(s, " ")
	c, _ := strconv.ParseUint(d[0], 10, 0)
	return uint(c), d[1]
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
