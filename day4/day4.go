package main

import (
	"bufio"
	"fmt"
	"os"
)

const file = "input.txt"

func main() {
	count := 0
	depot := readInput()

	for {
		changed := false
		for i, line := range depot {
			for j, val := range line {
				if val == 1 {
					if countTouching(depot, i, j) < 4 {
						depot[i][j] = 0
						count++
						changed = true
					}
				}
			}
		}
		if !changed {
			break
		}
	}
	fmt.Println(count)
}

func readInput() [][]int {
	file, err := os.Open(file)
	if err != nil {
		return nil
	}
	defer file.Close()
	var depot [][]int
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	// scan each line and check each character
	for scanner.Scan() {
		var isle []int
		for _, char := range scanner.Text() {
			if char == '@' {
				isle = append(isle, 1)
			} else {
				isle = append(isle, 0)
			}
		}
		depot = append(depot, isle)
	}
	return depot
}

func countTouching(depot [][]int, row, col int) int {
	directions := [][2]int{
		{-1, 0}, {-1, 1}, {0, 1}, {1, 1},
		{1, 0}, {1, -1}, {0, -1}, {-1, -1},
	}

	touching := 0

	for _, dir := range directions {
		neighborRow := row + dir[0]
		neighborCol := col + dir[1]
		if neighborRow >= 0 && neighborRow < len(depot) && neighborCol >= 0 && neighborCol < len(depot[neighborRow]) {
			if depot[neighborRow][neighborCol] == 1 {
				touching++
			}
		}
	}
	return touching
}
