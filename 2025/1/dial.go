package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Dial struct {
	current int
	max     int
}

func NewDial() *Dial {
	return &Dial{50, 99}
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

	if dir == "L" {
		steps *= -1
	}
	dial.current = (dial.current + steps) % (dial.max + 1)
	return dial.current
}
