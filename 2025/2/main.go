package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func IsRepeatedPattern(s string) bool {
	doubled := s + s
	return strings.Contains(doubled[1:len(doubled)-1], s)
}

// Check how many invalid ids are there in the range
// and return their sum
func AnalyzeRange(rng string) int {
	sum := 0
	parts := strings.SplitN(rng, "-", 2)

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(fmt.Sprintf("unable to parse start: %v\n%s", parts[0], err))
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(fmt.Sprintf("unable to parse end: %v\n%s", parts[1], err))
	}

	for num := start; num <= end; num++ {
		sNum := strconv.Itoa(num)
		if len(sNum) == 1 {
			// single digit cannot repeat
			continue
		}
		// in double digit numbers only multiples of 11 are repeated pattern
		if len(sNum) == 2 && num%11 != 0 {
			continue
		}
		if len(sNum) == 3 && num%111 != 0 {
			continue
		}

		// part one solution
		// if len(sNum)%2 != 0 {
		// 	// only even len numbers can consist of repeated sequence
		// 	continue
		// }
		// mid := len(sNum) / 2
		// seq := sNum[:mid]
		// rest := sNum[mid:]
		// if seq == rest {
		// 	fmt.Printf("ID INVALID: invalid id: %s\n", sNum)
		// 	sum += num
		// }

		for seqEndIdx := 1; seqEndIdx <= len(sNum)/2; seqEndIdx++ {
			isInvalid := true
			seq := sNum[:seqEndIdx]
			rest := sNum[seqEndIdx:]

			// no need to check the pattern if string is not dividable in these patterns
			if len(rest)%len(seq) != 0 {
				isInvalid = false
				continue
			}

			for i := 0; i < len(rest); i += len(seq) {
				if seq != rest[i:i+len(seq)] {
					isInvalid = false
					break
				} else {
				}
			}

			if isInvalid {
				// fmt.Printf("ID INVALID: invalid id: %s\n", sNum)
				sum += num
				break
			}
		}
	}

	return sum
}

func Solve(input string) int {
	solution := 0
	for value := range strings.SplitSeq(input, ",") {
		solution += AnalyzeRange(strings.Trim(value, " \n"))
	}
	return solution
}

func AnalyzeQuicker(rng string) int {
	sum := 0
	parts := strings.SplitN(rng, "-", 2)

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(fmt.Sprintf("unable to parse start: %v\n%s", parts[0], err))
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(fmt.Sprintf("unable to parse end: %v\n%s", parts[1], err))
	}

	for num := start; num <= end; num++ {
		sNum := strconv.Itoa(num)
		if len(sNum) == 1 {
			// single digit cannot repeat
			continue
		}
		// in double digit numbers only multiples of 11 are repeated pattern
		if len(sNum) == 2 && num%11 != 0 {
			continue
		}
		if len(sNum) == 3 && num%111 != 0 {
			continue
		}
		if IsRepeatedPattern(sNum) {
			sum += num
		}
	}
	return sum
}

func SolveQuicker(input string) int {
	solution := 0
	for value := range strings.SplitSeq(input, ",") {
		solution += AnalyzeQuicker(strings.Trim(value, " \n"))
	}
	return solution
}

// Quicker solution did not turn out to be quicker,
// I think this is because it still performs linear contains operation inside 2 nested loops
// and strings we are checking are relatively small
func main() {
	// TODO: reading thing byte by byte would be cool and efficient
	byteContent, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Sprintf("Unable to read file: %s", err))
	}
	content := string(byteContent)

	start := time.Now()
	solution := Solve(content)
	elapsed := time.Since(start)
	fmt.Printf("The solution is %d; Took: %s\n", solution, elapsed)

	start = time.Now()
	solution = SolveQuicker(content)
	elapsed = time.Since(start)
	fmt.Printf("The solution is %d; Took: %s\n", solution, elapsed)
}
