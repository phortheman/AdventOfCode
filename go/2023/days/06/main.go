package main

import (
	file "aoc23/internal"
	"fmt"
	"math"
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
			localRange := QuadraticFormula(time, distance)
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

	return QuadraticFormula(trueTime, trueDistance)
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

// a = 1 (positive parabola), b = time, , c = min distance
func QuadraticFormula(time, distance int) int {
	// Distance + 1 because we need to beat the record
	discriminant := time*time - 4*(distance+1)

	// Cache this math so it is only done once
	sqrtDiscriminant := math.Sqrt(float64(discriminant))

	// lTime needs the ceiling of the float. lTime = 10.1111 would mean 10.000 is too low
	lTime := int(math.Ceil(math.Abs((float64(-time) + sqrtDiscriminant) / 2)))

	// rTime needs the floor of the float. rTime = 10.8888 would mean 11.000 is too high
	rTime := int(math.Floor(math.Abs((float64(-time) - sqrtDiscriminant) / 2)))

	// +1 to make it inclusive
	return rTime - lTime + 1
}
