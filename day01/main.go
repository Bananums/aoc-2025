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

	fmt.Println("Puzzle Largest number:", FindSetLargest(lines))

	fmt.Println("Part 1:", solvePart1(lines, false))
	fmt.Println("Part 2:", solvePart2(lines, false))
}

func solvePart1(lines []string, verbose bool) int {
	var dial = 50
	var hits = 0
	var rollover = 100

	for _, line := range lines {
		firstChar := string(line[0])
		numberStr := line[1:]
		number, _ := strconv.Atoi(numberStr)

		if number >= rollover {
			// Does not handle numbers above 999. No point in doing more, since largest puzzle number is 997
			number = number % 100
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

func solvePart2(lines []string, verbose bool) int {
	var dial = 50
	var hits = 0
	var rollover = 100

	for _, line := range lines {
		firstChar := string(line[0])
		numberStr := line[1:]
		number, _ := strconv.Atoi(numberStr)

		for number > 100 {
			hits++
			number -= 100
		}

		dialOld := dial
		if firstChar == "R" {
			if (dial + number) >= rollover {
				dial = (dial + number) - rollover
				if dial != 0 && dialOld != 0 {
					hits++
				}
			} else {
				dial += number
			}
		} else {
			if (dial - number) < 0 {
				dial = rollover - (number - dial)
				if dial != 0 && dialOld != 0 {
					hits++
				}
			} else {
				dial -= number
			}
		}

		if dial == 0 {
			hits++
		}

		if verbose {
			fmt.Println(line, "- dial:", dial, "hits:", hits)
		}

	}

	return hits

}

func FindSetLargest(lines []string) int {
	var number = 0
	for _, line := range lines {
		locNumberStr := line[1:]
		locNumber, _ := strconv.Atoi(locNumberStr)
		if locNumber > number {
			number = locNumber
		}
	}
	return number
}
