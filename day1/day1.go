package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var FILE string = "input.txt"

type instruction struct {
	direction rune
	value     int
}

func main() {
	dialHead := 50
	zeroCounter := 0
	for _, instruction := range getInstructions() {
		result := rotateDial(dialHead, instruction)
		dialHead = result[0]
		zeroCounter += result[1]
	}
	fmt.Printf("Zero Counter: %d\n", zeroCounter)
}

func getInstructions() []instruction {
	file, err := os.Open(FILE)
	if err != nil {
		panic(err)
	}
	var instructions []instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		direction := rune(line[0])
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, instruction{direction, value})
	}
	defer file.Close()
	return instructions
}

func rotateDial(dialHead int, inst instruction) []int {
	zeroCounter := 0
	valCount := inst.value
	if inst.direction == 'R' {
		for valCount > 0 {
			dialHead = (dialHead + 1) % 100
			if dialHead == 0 {
				zeroCounter++
			}
			valCount--
		}
	} else if inst.direction == 'L' {
		for valCount > 0 {
			dialHead = (dialHead - 1 + 100) % 100
			if dialHead == 0 {
				zeroCounter++
			}
			valCount--
		}
	}
	return []int{dialHead, zeroCounter}
}
