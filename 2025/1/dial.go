package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Dial struct {
	current  int
	max      int
	Solution int
}

func NewDial() *Dial {
	return &Dial{50, 99, 0}
}

func (dial *Dial) Turn(input string) int {
	if len(input) < 2 {
		panic(fmt.Errorf("invalid input: %v", input))
	}
	input = strings.ToUpper(input)
	dir := string(input[0])

	steps, err := strconv.Atoi(input[1:])
	if err != nil {
		panic(fmt.Errorf("unable to parse int: %s", input[1:]))
	}
	/*
		 * current 31
		 * L164
			* reaches 0 still 133 to L
			* reaches 0 still 33 to L
			* current points to 66
			*
			*
			* Solution:
			* 1) // 100 => how many times reaches 0 for sure
			* 2) take the reminder of % 100
			* 3) if l then current < reminder => reaches zero
			* 4) if r then current + reminder > 99 reaches zero
	*/

	dial.Solution += steps / 100
	steps = steps % 100

	if dir == "L" {
		if dial.current != 0 && dial.current <= steps {
			dial.Solution++
		}
		steps *= -1
	} else if dial.current+steps > dial.max {
		dial.Solution++
	}

	dial.current = (dial.current + steps) % (dial.max + 1)
	if dial.current < 0 {
		dial.current = dial.max + 1 + dial.current
	}
	return dial.current
}
