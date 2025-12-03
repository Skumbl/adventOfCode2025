package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const file = "input.txt"

func main() {
	total := 0
	for _, batteryBank := range readInput() {
		val := findLargestJoltage(batteryBank, 0, 11)
		fmt.Println(val)
		total += (val)
	}
	fmt.Println(total)
}

func readInput() [][]int {
	file, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var allBatteryBanks [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var batteryBank []int
		line := scanner.Text()
		for _, char := range line {
			digit := int(char - '0')
			batteryBank = append(batteryBank, digit)
		}
		allBatteryBanks = append(allBatteryBanks, batteryBank)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return allBatteryBanks
}

func findLargestJoltage(batteryBank []int, total int, loop int) int {
	if loop < 0 || len(batteryBank) <= loop {
		return total
	}
	largestValue, index := 0, 0
	for i, digit := range batteryBank[:len(batteryBank)-loop] {
		if digit > largestValue {
			largestValue = digit
			index = i
		}
	}
	total += largestValue * int(math.Pow(10, float64(loop)))
	return findLargestJoltage(batteryBank[index+1:], total, loop-1)
}
