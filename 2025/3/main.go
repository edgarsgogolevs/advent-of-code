package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findMaxJoltagePart1(bank string) int {
	max := [2]int{0, 0}
	for i, c := range bank {
		val := int(c - '0')
		if val > max[1] {
			max[1] = val
		}
		if i == len(bank)-1 {
			break
		}
		if val > max[0] {
			max[0] = val
			max[1] = 0
		}
	}

	result, _ := strconv.Atoi(strconv.Itoa(max[0]) + strconv.Itoa(max[1]))
	return result
}

func findMaxJoltage(bank string) int {
	// max := [12]int{}
	max := 0
	maxIdx := -1
	result := ""
	for len(result) < 12 {
		for i := maxIdx + 1; i < len(bank)-12+len(result)+1; i++ {
			val := int(bank[i] - '0')
			if val > max {
				max = val
				maxIdx = i
			}
		}
		result += strconv.Itoa(max)
		max = 0
	}
	resultInt, _ := strconv.Atoi(result)
	return resultInt
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic("Could not open the file")
	}
	scanner := bufio.NewScanner(f)

	result := 0

	for scanner.Scan() {
		// battery bank
		bank := scanner.Text()
		if len(bank) == 0 {
			continue
		}
		// bankMaxJoltage := findMaxJoltagePart1(bank)
		bankMaxJoltage := findMaxJoltage(bank)
		result += bankMaxJoltage
		fmt.Printf("Max joltage of the bank: %d\n", bankMaxJoltage)
	}
	fmt.Printf("Result: %d\n", result)
}
