package main

import (
	file "aoc23/internal"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var EXAMPLE string = `Time:      7  15   30
Distance:  9  40  200`

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

	timeData := ParseData(content[0])
	distanceData := ParseData(content[1])

	if len(timeData) != len(distanceData) {
		fmt.Printf("Data mis-match. Time len: %v | Distance len: %v\n", len(timeData), len(distanceData))
		os.Exit(1)
	}

	var part1Total int = Part1Solver(timeData, distanceData)
	var part2Total int = Part2Solver(timeData, distanceData)

	fmt.Println(part1Total)
	fmt.Println(part2Total)
}

func Part1Solver(timeData, distanceData []int) int {
	var total int = 1
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i := range timeData {
		wg.Add(1)
		go func(time, distance int) {
			defer wg.Done()
			localRange := ProcessData(time, distance)
			mutex.Lock()
			total *= localRange
			mutex.Unlock()
		}(timeData[i], distanceData[i])
	}
	wg.Wait()
	return total
}

func Part2Solver(timeData, distanceData []int) int {
	var temp string
	for _, d := range timeData {
		temp += fmt.Sprint(d)
	}
	trueTime, _ := strconv.Atoi(temp)
	temp = ""
	for _, d := range distanceData {
		temp += fmt.Sprint(d)
	}
	trueDistance, _ := strconv.Atoi(temp)

	return ProcessData(trueTime, trueDistance)
}

func GetDistance(timeHeld, timeEnd int) int {
	if timeHeld >= timeEnd {
		return 0
	}
	return timeHeld * (timeEnd - timeHeld)
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func BuildIntFromByte(c byte, v int) int {
	return v*10 + int(c) - '0'
}

func ParseData(line []byte) []int {
	var data []int
	for i := 0; i < len(line); {
		if IsDigit(line[i]) {
			t := 0
			for IsDigit(line[i]) {
				t = BuildIntFromByte(line[i], t)
				i++
				if i >= len(line) {
					break
				}
			}
			data = append(data, t)
		} else {
			i++
		}
	}
	return data
}

/*
Find the first millisecond that breaks the record and then keep increasing
until the time no longer breaks the record and stop processing
*/
func ProcessData(time, distance int) int {
	lower, upper := -1, -1
	for j := 0; j <= time; j++ {
		if lower < j && GetDistance(j, time) > distance && lower == -1 {
			lower = j
		} else if upper < j && GetDistance(j, time) <= distance && lower != -1 {
			upper = j
			break
		}
	}
	return upper - lower
}
