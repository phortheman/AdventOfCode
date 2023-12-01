package cmd

import (
	"aoc23/cmd"
	file "aoc23/internal"
	"testing"
)

func TestDayOnePartOne(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	content := file.Read_String_Into_Byte_Slice(input)
	var total int
	var expected int = 142
	for _, line := range content {
		total += cmd.DayOnePartOne(line)
	}
	if total != expected {
		t.Errorf("Expected %d and got %d", expected, total)
	}
}
