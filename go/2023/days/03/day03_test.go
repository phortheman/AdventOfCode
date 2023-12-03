package main

import (
	file "aoc23/internal"
	"testing"
)

var example string = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

var example2 string = `...155...............259..........629.646..............*194...111...............%...................998.....292................872..........
.......444.......648.....730.154....*.....657..$....816........*.....%.......618..123.....-...........*.......@........772.@....+.......*200
..........*.........*.....&....*............*..208......168#..305....510...........*.....907.........901.........582.....$.238........50....`

func TestDayThreePartOne(t *testing.T) {
	input := file.Read_String_Into_Byte_Slice(example)
	total := Part1Solver(input)
	var expected uint = 4361
	if total != expected {
		t.Errorf("Expected %d and got %d", expected, total)
	}
}
func TestDayThreePartTwo(t *testing.T) {
	input := file.Read_String_Into_Byte_Slice(example)
	total := Part2Solver(input)
	var expected uint = 467835
	if total != expected {
		t.Errorf("Expected %d and got %d", expected, total)
	}
}
func TestDayThreePartTwoExample2(t *testing.T) {
	input := file.Read_String_Into_Byte_Slice(example2)
	total := Part2Solver(input)
	var expected uint = 1_101_357
	if total != expected {
		t.Errorf("Expected %d and got %d", expected, total)
	}
}
