package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const relative_input string = "../../../inputs/2015/04/input.txt"

var bTest = false
var inputFileName string

func init() {
	flag.StringVar(&inputFileName, "i", "",
		"Path to the puzzle input. "+
			"Default to using the internal relative path. "+
			"Pass 'stdin' to use it instead")
}

func main() {
	flag.Parse()
	var content []byte
	var err error
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

	start := time.Now()
	part1, part2 := Solver(string(content))
	duration := time.Since(start)

	fmt.Printf("Time: %v\nPart 1: %d\nPart 2: %d\n", duration, part1, part2)
}

func Solver(input string) (int, int) {
	// For some reason os.ReadFile puts a line feed in the byte slice at the end ???
	secret := strings.TrimSpace(input)
	var part1, part2 int
	var checkValue int = 1
	h := md5.New()
	var sum []byte
	for {
		io.WriteString(h, secret)
		io.WriteString(h, strconv.Itoa(checkValue))
		sum = h.Sum(nil)

		checkSum := fmt.Sprintf("%x", sum)
		if part1 == 0 && strings.HasPrefix(checkSum, "00000") {
			part1 = checkValue
		}
		if part2 == 0 && strings.HasPrefix(checkSum, "000000") {
			part2 = checkValue
		}

		if part1 != 0 && part2 != 0 {
			break
		}

		checkValue++
		clear(sum)
		h.Reset()
	}

	return part1, part2
}
