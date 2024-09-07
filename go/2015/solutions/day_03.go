package solutions

import "fmt"

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

func (p *Point) AddPoint(h *map[string]int) {
	(*h)[p.GetPosition()] += 1
}

func Day03Solver(input []byte) (int, int) {
	part1 := make(map[string]int)
	pointPart1 := NewPoint()
	pointPart1.AddPoint(&part1)

	part2 := make(map[string]int)
	pointsPart2 := [2]Point{NewPoint(), NewPoint()}
	pointsPart2[0].AddPoint(&part2)
	pointsPart2[1].AddPoint(&part2)

	for idx, char := range input {
		pointPart1.ChangePoint(char)
		pointPart1.AddPoint(&part1)
		pointsPart2[idx%2].ChangePoint(char)
		pointsPart2[idx%2].AddPoint(&part2)
	}

	return len(part1), len(part2)
}
