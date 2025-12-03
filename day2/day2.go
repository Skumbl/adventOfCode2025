package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var FILE string = "input.txt"

type Range struct {
	start int64
	end   int64
}

func main() {
	ranges := readInput()
	var totaledInvalid int64 = 0
	for _, r := range ranges {
		invalidNumbers := invalidRanges(r)
		for _, num := range invalidNumbers {
			totaledInvalid += num
		}
	}
	fmt.Println(totaledInvalid)
}

func readInput() []Range {
	file, err := os.Open(FILE)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	var ranges []Range
	parts := strings.Split(line, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		rangeParts := strings.Split(part, "-")
		start, _ := strconv.ParseInt(rangeParts[0], 10, 64)
		end, _ := strconv.ParseInt(rangeParts[1], 10, 64)
		ranges = append(ranges, Range{start, end})
	}
	return ranges
}

func invalidRanges(r Range) []int64 {
	var invalidNumbers []int64
	for i := r.start; i <= r.end; i++ {
		if isInvalid(i) {
			invalidNumbers = append(invalidNumbers, i)
		}
	}
	return invalidNumbers
}

func isInvalid(num int64) bool {
	str := strconv.FormatInt(num, 10)
	length := len(str)
	for blocksize := 1; blocksize < length/2+1; blocksize++ {
		if length%blocksize != 0 {
			continue
		}
		block := str[0:blocksize]
		repeating := true
		for pos := blocksize; pos < length; pos += blocksize {
			currBlock := str[pos : pos+blocksize]
			if currBlock != block {
				repeating = false
				break
			}
		}
		if repeating {
			return true
		}
	}
	return false
}
