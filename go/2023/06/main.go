package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"sync"
	"time"
)

const relative_input string = "../../../inputs/2023/06/input.txt"
const sampleInput = `Time:      7  15   30
Distance:  9  40  200`

var bTest = false
var inputFileName string

func init() {
	flag.BoolVar(&bTest, "t", false, "Run using the sample input")
	flag.StringVar(&inputFileName, "i", "",
		"Path to the puzzle input. "+
			"Default to using the internal relative path. "+
			"Pass 'stdin' to use it instead")
}

func main() {
	flag.Parse()
	var content []byte = []byte(sampleInput)
	var err error
	if !bTest {
		switch inputFileName {
		case "stdin":
			content, err = io.ReadAll(os.Stdin)
		case "":
			content, err = os.ReadFile(relative_input)
		default:
			content, err = os.ReadFile(inputFileName)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to read file: %v", err)
			os.Exit(66)
		}
	}

	start := time.Now()
	part1, part2 := Solver(content)
	duration := time.Since(start)

	fmt.Printf("Time: %v\nPart 1: %d\nPart 2: %d\n", duration, part1, part2)
}

func Solver(input []byte) (int, int) {
	data := bytes.Split(input, []byte("\n"))
	timeData := ParseData(data[0])
	distanceData := ParseData(data[1])

	var part1Total int = Part1(timeData, distanceData)
	var part2Total int = Part2(timeData, distanceData)
	return part1Total, part2Total
}

func Part1(timeData, distanceData []int) int {
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

func Part2(timeData, distanceData []int) int {
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
