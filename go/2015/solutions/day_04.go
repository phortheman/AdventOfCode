package solutions

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Day04Solver(input string) (int, int) {
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
