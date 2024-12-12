package main

import (
	file "aoc23/internal"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var EXAMPLE string = `seeds: 79 14 55 13

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

	seeds, data := ParseData(content)
	var part1Total int = Part1Solver(seeds, data)
	var part2Total int = Part2Solver(seeds, data)

	fmt.Println(part1Total)
	fmt.Println(part2Total)
}

func Part1Solver(seeds []int, data []SourceDestination) int {
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
func Part2Solver(seeds []int, data []SourceDestination) int {
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
