package main

import (
	"adventOfCode2025/file"
	"fmt"
	"strconv"
)

func main() {
	lines := readInput("input.txt")
	maxWidth := maxLen(lines)
	operatorRow := len(lines) - 1

	grandTotal := 0

	col := maxWidth - 1
	for col >= 0 {
		for col >= 0 && isSpaceColumn(lines, col, operatorRow) {
			col--
		}
		if col < 0 {
			break
		}
		var numbers []int
		var operator byte

		for col >= 0 && !isSpaceColumn(lines, col, operatorRow) {
			numStr := ""
			for row := 0; row < operatorRow; row++ {
				if col < len(lines[row]) && lines[row][col] >= '0' && lines[row][col] <= '9' {
					numStr += string(lines[row][col])
				}
			}
			if numStr != "" {
				num, _ := strconv.Atoi(numStr)
				numbers = append(numbers, num)
			}
			if col < len(lines[operatorRow]) {
				char := lines[operatorRow][col]
				if char == '+' || char == '*' {
					operator = char
				}
			}

			col--
		}

		if len(numbers) > 0 && operator != 0 {
			result := numbers[0]
			for i := 1; i < len(numbers); i++ {
				if operator == '+' {
					result += numbers[i]
				} else if operator == '*' {
					result *= numbers[i]
				}
			}
			fmt.Printf("Problem: %v %c = %d\n", numbers, operator, result)
			grandTotal += result
		}
	}
	fmt.Printf("\nGrand Total: %d\n", grandTotal)
}

func readInput(filename string) []string {
	grid := file.Read2DArray(filename, "")
	var lines []string
	for _, row := range grid {
		line := ""
		for _, char := range row {
			line += char
		}
		lines = append(lines, line)
	}
	return lines
}

func maxLen(lines []string) int {
	max := 0
	for _, line := range lines {
		if len(line) > max {
			max = len(line)
		}
	}
	return max
}

func isSpaceColumn(lines []string, col int, operatorRow int) bool {
	for row := 0; row <= operatorRow; row++ {
		if col < len(lines[row]) && lines[row][col] != ' ' {
			return false
		}
	}
	return true
}
