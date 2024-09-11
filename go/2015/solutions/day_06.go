package solutions

import (
	"bufio"
	"strconv"
	"strings"
)

type gridState [][]bool

func (l gridState) onLights() int {
	var on int
	for i := range l {
		for j := range l[i] {
			if l[i][j] {
				on++
			}
		}
	}
	return on
}

type gridBrightness [][]int

func (l gridBrightness) brightness() int {
	var brightness int
	for i := range l {
		for j := range l[i] {
			brightness += l[i][j]
		}
	}
	return brightness

}

var part1Grid gridState = make(gridState, 1_000)
var part2Grid gridBrightness = make(gridBrightness, 1_000)

func Day06Solver(input string) (int, int) {
	var part1, part2 int
	for i := range part1Grid {
		part1Grid[i] = make([]bool, 1_000)
		part2Grid[i] = make([]int, 1_000)
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		if tokens[0] == "toggle" {
			togglePart1Grid(tokens[1], tokens[3])
			statePart2Grid(tokens[1], tokens[3], 2)
		} else {
			statePart1Grid(tokens[2], tokens[4], (tokens[1] == "on"))
			val := -1
			if tokens[1] == "on" {
				val = 1
			}
			statePart2Grid(tokens[2], tokens[4], val)
		}
	}

	part1 = part1Grid.onLights()
	part2 = part2Grid.brightness()

	return part1, part2
}

func togglePart1Grid(start, end string) {
	startX, startY := getRange(start)
	endX, endY := getRange(end)

	for i := startY; i <= endY; i++ {
		for j := startX; j <= endX; j++ {
			part1Grid[i][j] = !part1Grid[i][j]
		}
	}
}

func statePart1Grid(start, end string, state bool) {
	startX, startY := getRange(start)
	endX, endY := getRange(end)

	for i := startY; i <= endY; i++ {
		for j := startX; j <= endX; j++ {
			part1Grid[i][j] = state
		}
	}
}

func statePart2Grid(start, end string, inc int) {
	startX, startY := getRange(start)
	endX, endY := getRange(end)

	for i := startY; i <= endY; i++ {
		for j := startX; j <= endX; j++ {
			if part2Grid[i][j] == 0 && inc < 0 {
				continue
			}
			part2Grid[i][j] += inc
		}
	}
}

func getRange(input string) (int, int) {
	val := strings.Split(input, ",")
	x, _ := strconv.Atoi(val[0])
	y, _ := strconv.Atoi(val[1])
	return x, y
}
