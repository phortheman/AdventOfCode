package day02

import (
	"bufio"
	"slices"
	"strconv"
	"strings"
)

func Solver(input string) (int, int) {
	part1 := 0
	part2 := 0

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		val := strings.Split(line, "x")
		rec := newRectangle(val[0], val[1], val[2])
		part1 += rec.getDimension()
		part2 += rec.getRibbon()
	}

	return part1, part2
}

type rectangle struct {
	length int
	width  int
	height int
}

func newRectangle(length, width, height string) rectangle {
	rectangle := rectangle{}
	rectangle.length, _ = strconv.Atoi(length)
	rectangle.width, _ = strconv.Atoi(width)
	rectangle.height, _ = strconv.Atoi(height)

	return rectangle
}

// Surface area of the rectangle plus some slack for the wrapping paper
func (r *rectangle) getDimension() int {
	var sum int
	dimensions := []int{
		r.length * r.width,
		r.width * r.height,
		r.height * r.length,
	}

	slack := dimensions[0]

	for _, dim := range dimensions {
		if slack > dim {
			slack = dim
		}
		sum += 2 * dim
	}

	return sum + slack
}

// Find the shortest distance to wrap the ribbon around then add l*w*h
func (r *rectangle) getRibbon() int {
	sum := r.length * r.width * r.height
	distances := []int{
		2*r.length + 2*r.width,
		2*r.width + 2*r.height,
		2*r.height + 2*r.length,
	}
	return sum + slices.Min(distances)
}
