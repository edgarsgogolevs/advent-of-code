package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve2() {
	f, err := os.Open("input")
	if err != nil {
		fmt.Println("unable to open input file")
		return
	}
	defer f.Close()

	dial := NewDial()

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		dial.Turn(line)
	}
	fmt.Printf("Solution:\t%d\n", dial.Solution)
}

func solve1() {
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
	dial.Turn("R350")
	fmt.Printf("Test solution: %d\n", dial.Solution)
}

func main() {
	solve1()
	solve2()
}
