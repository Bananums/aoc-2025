package main

import (
	"aoc-2025/internal/util"
	"embed"
	"fmt"
	"log"
	"strconv"
)

//go:embed example.txt puzzle.txt
var inputs embed.FS

func main() {
	fmt.Println("Advent of Code - Day 01")

	var lines []string
	var err error
	lines, err = util.LoadFile("puzzle.txt", inputs)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", solvePart1(lines, false))
	fmt.Println("Part 2:", solvePart2(lines))
}

func solvePart1(lines []string, verbose bool) int {
	var dial int = 50
	var hits int = 0
	var rollover int = 100

	for _, line := range lines {
		firstChar := string(line[0])
		numberStr := line[1:]
		number, _ := strconv.Atoi(numberStr)

		if number >= rollover {
			number = number % 100 // Does not handle numbers above 1000
		}

		if firstChar == "R" {
			if (dial + number) >= rollover {
				dial = (dial + number) - rollover
			} else {
				dial += number
			}
		} else {
			if (dial - number) < 0 {
				dial = rollover - (number - dial)
			} else {
				dial -= number
			}
		}

		if verbose {
			fmt.Println(line, "-", firstChar, "-", number, " - dial:", dial)
		}

		if dial == 0 {
			hits++
		}
	}

	return hits
}

func solvePart2(lines []string) int {
	return 0
}
