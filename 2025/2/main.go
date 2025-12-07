package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Check how many invalid ids are there in the range
// and return their sum
func AnalyzeRange(rng string) int {
	// TODO: must be some neat optimization, that reduces nested loop (currently 4 O_o)
	//
	// TODO: possible optimization:
	// as number should consist only of repeating sequences,
	// we first longest sequence of not repeating char, for example,
	// in 123123123 it would be 123
	// then check if rest of string is repeating the crap
	//
	// in 123123123
	// take 1
	// check if next seq of same len is the same
	// 1 != 2
	// take 12
	// 12 != 31
	// take 123
	// 123 == 123
	// start with last index + 1 and take slices of len of seq until the end
	// in this case take the last 123
	// 123 == 123 => invalid ID
	//
	// TODO: there is also a trick of doubling the string getting rid of first and last char
	// and checking if original string is a substring of this one, if is => repeating seq
	//
	// TODO: implement the trick and compare times of solutions
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

			for i := 0; i < len(rest); i += len(seq) {
				if i+len(seq) > len(rest) {
					isInvalid = false
					break
				}
				if seq != rest[i:i+len(seq)] {
					isInvalid = false
					break
				} else {
				}
			}

			if isInvalid {
				fmt.Printf("ID INVALID: invalid id: %s\n", sNum)
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

func main() {
	// TODO: reading thing byte by byte would be cool and efficient
	byteContent, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Sprintf("Unable to read file: %e", err))
	}
	solution := Solve(string(byteContent))

	fmt.Printf("The solution is %d\n", solution)
}
