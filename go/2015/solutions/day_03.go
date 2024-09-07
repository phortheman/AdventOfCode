package solutions

import "fmt"

func Day03Solver(input []byte) (int, int) {
	part1 := make(Grid)
	pointPart1 := NewPoint()
	part1.AddPoint(pointPart1)

	part2 := make(Grid)
	pointsPart2 := [2]Point{NewPoint(), NewPoint()}
	part2.AddPoint(pointsPart2[0])
	part2.AddPoint(pointsPart2[1])

	for idx, char := range input {
		pointPart1.ChangePoint(char)
		part1.AddPoint(pointPart1)
		pointsPart2[idx%2].ChangePoint(char)
		part2.AddPoint(pointsPart2[idx%2])
	}

	return len(part1), len(part2)
}

// Key is "x,y". Value is the number of times that point was passed
type Grid map[string]int

func (g *Grid) AddPoint(p Point) {
	(*g)[p.GetPosition()] += 1
}

type Point struct {
	x int
	y int
}

func NewPoint() Point {
	return Point{0, 0}
}

func (p *Point) ChangePoint(instruction byte) {
	switch instruction {
	case '<':
		p.x -= 1
	case '>':
		p.x += 1
	case '^':
		p.y += 1
	case 'v':
		p.y -= 1
	}
}

func (p *Point) GetPosition() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}
