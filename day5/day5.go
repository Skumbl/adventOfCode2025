package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const file = "input.txt"

type numRange struct {
	start int
	end   int
}

// part 2 of advent of code, massively changed from part one
func main() {
	ranges, ids := readInput()
	fmt.Println(countValidIds(ranges, ids))
}

func readInput() ([]numRange, []int) {
	var ranges []numRange
	var ids []int
	readRanges := true

	file, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
			parts := strings.Split(line, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			fmt.Println(start, end)
			ranges = append(ranges, numRange{start, end})
		} else {
			parts := strings.Split(line, " ")
			number, _ := strconv.Atoi(parts[0])
			fmt.Println(number)
			ids = append(ids, number)
		}
	}

	return ranges, ids
}

func countValidIds(ranges []numRange, ids []int) int {
	count := 0
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r.start && id <= r.end {
				count++
				break
			}
		}
	}
	return count
}
