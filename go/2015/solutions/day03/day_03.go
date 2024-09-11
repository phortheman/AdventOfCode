package day03

import "fmt"

func Solver(input []byte) (int, int) {
	part1 := make(grid)
	pointPart1 := newPoint()
	part1.addPoint(pointPart1)

	part2 := make(grid)
	pointsPart2 := [2]point{newPoint(), newPoint()}
	part2.addPoint(pointsPart2[0])
	part2.addPoint(pointsPart2[1])

	for idx, char := range input {
		pointPart1.changePoint(char)
		part1.addPoint(pointPart1)
		pointsPart2[idx%2].changePoint(char)
		part2.addPoint(pointsPart2[idx%2])
	}

	return len(part1), len(part2)
}

// Key is "x,y". Value is the number of times that point was passed
type grid map[string]int

func (g *grid) addPoint(p point) {
	(*g)[p.getPosition()] += 1
}

type point struct {
	x int
	y int
}

func newPoint() point {
	return point{0, 0}
}

func (p *point) changePoint(instruction byte) {
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

func (p *point) getPosition() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}
