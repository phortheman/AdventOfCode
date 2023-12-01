package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Read_File_Into_Memory(filePath string) ([][]byte, error) {
	// Load the file into memory
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Determine the number of lines in the file as that is more likey to be a
	// massive number to reduce copying and reallocating massive slices
	lineCount := strings.Count(string(content), "\n") + 1

	lines := make([][]byte, lineCount)
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	for i := 0; scanner.Scan(); i++ {
		// Allocate the exact line length before copying the bytes over.
		// The scanner.Bytes() has a capacity of 4096 so this will reduce the memory usage for shorter lines
		lines[i] = make([]byte, len(scanner.Bytes()))
		copy(lines[i], scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file: ", err)
		return nil, err
	}

	return lines, nil
}

func Read_String_Into_Byte_Slice(input string) [][]byte {
	content := strings.Split(input, "\n")
	lines := make([][]byte, len(content))
	for i, line := range content {
		lines[i] = make([]byte, len(line))
		copy(lines[i], line)
	}
	return lines
}
