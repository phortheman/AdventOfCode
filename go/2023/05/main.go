package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const relative_input string = "../../../inputs/2023/05/input.txt"
const sampleInput = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

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
	rawData := bytes.Split(input, []byte("\n"))
	seeds, data := ParseData(rawData)
	var part1 int = Part1(seeds, data)
	var part2 int = Part2(seeds, data)

	return part1, part2
}

func Part1(seeds []int, data []SourceDestination) int {
	var location int = -1
	for _, seed := range seeds {
		t := GetLocation(seed, 0, data)
		if t < location || location == -1 {
			location = t
		}
	}
	return location
}

// Horrible brute force solution. Need to research a better algorithm but tried using goroutines to lessen the pain
// 639s single core
// 293s multi core
func Part2(seeds []int, data []SourceDestination) int {
	var location = -1
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i := 0; i+1 < len(seeds); i += 2 {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			goRouteLocation := -1
			for j := start; j < start+end; j++ {
				t := GetLocation(j, 0, data)
				if t < goRouteLocation || goRouteLocation == -1 {
					goRouteLocation = t
				}
			}
			mutex.Lock()
			if goRouteLocation < location || location == -1 {
				location = goRouteLocation
			}
			mutex.Unlock()
		}(seeds[i], seeds[i+1])
	}
	wg.Wait()
	return location
}

func ParseData(input [][]byte) ([]int, []SourceDestination) {
	var seeds []int
	var data []SourceDestination
	var parsing bool
	var curMap int = 0
	for i, line := range input {
		if i == 0 {
			seeds = append(seeds, SplitDataIntoInts(line[7:])...)
			continue
		}
		if i == 1 {
			continue
		}
		if len(line) == 0 {
			parsing = false
			curMap++
			continue
		}
		if ByteContains(line, "map:") {
			parsing = true
			data = append(data, SourceDestination{})
			continue
		}
		if parsing {
			var destination, source, length int
			d := SplitDataIntoInts(line)
			if len(d) == 3 {
				destination = d[0]
				source = d[1]
				length = d[2]
			} else {
				fmt.Println("Got unexpected int slice length. Line: ", string(line))
				fmt.Println("Length: ", len(d))
				os.Exit(1)
			}
			data[curMap].AddData(source, destination, length)
		}
	}
	return seeds, data
}

type SourceDestination struct {
	sourceStarts      []int
	destinationStarts []int
	lengths           []int
}

func (sd *SourceDestination) AddData(s, d, l int) {
	sd.sourceStarts = append(sd.sourceStarts, s)
	sd.destinationStarts = append(sd.destinationStarts, d)
	sd.lengths = append(sd.lengths, l)
}

func (sd *SourceDestination) GetDestination(s int) int {
	for i, n := range sd.sourceStarts {
		if s >= n && s < n+sd.lengths[i] {
			return sd.destinationStarts[i] + s - n
		}
	}
	return s
}

func GetLocation(source, i int, data []SourceDestination) int {
	if i == len(data) {
		return source
	}
	return GetLocation(data[i].GetDestination(source), i+1, data)
}

func SplitDataIntoInts(d []byte) []int {
	var output []int
	s := string(d)
	for _, n := range strings.Split(s, " ") {
		v, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println("Error splitting data into ints. Line: ", s)
			os.Exit(1)
		}
		output = append(output, v)
	}
	return output
}

func ByteContains(b []byte, s string) bool {
	sb := []byte(s)
	return bytes.Contains(b, sb)
}
