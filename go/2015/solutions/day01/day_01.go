package day01

func Solver(input []byte) (int, int) {
	part1 := 0
	part2 := 0
	inBasement := false

	for _, char := range input {
		if !inBasement {
			part2 += 1
		}
		if char == '(' {
			part1 += 1
		} else if char == ')' {
			part1 -= 1
		}
		if part1 < 0 {
			inBasement = true
		}
	}
	return part1, part2
}
