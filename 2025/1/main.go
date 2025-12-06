package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve() {
	f, err := os.Open("input")
	if err != nil {
		fmt.Println("Unable to read input file")
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	dial := NewDial()
	solution := 0

	for scanner.Scan() {
		line := scanner.Text()
		if dial.Turn(line) == 0 {
			solution++
		}
	}
	fmt.Printf("Solution:\t%d\n", solution)
}

func test() {
	dial := NewDial()
	fmt.Printf("%v\n", dial.Turn("L100"))
}

func main() {
	solve()
}
