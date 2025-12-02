package main

import (
	"aoc-2025/internal/util"
	"embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed example.txt puzzle.txt
var inputs embed.FS

func main() {
	fmt.Println("Advent of Code - Day 02")

	lines, err := util.LoadCommaFile("puzzle.txt", inputs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", solvePart1(lines))
	fmt.Println("Part 2:", solvePart2(lines))
}

func solvePart1(lines []string) int {

	sum := 0

	for _, line := range lines {
		fmt.Println(line)

		rangeSpanStr := strings.Split(line, "-")
		var rangeSpan [2]int
		rangeSpan[0], _ = strconv.Atoi(rangeSpanStr[0])
		rangeSpan[1], _ = strconv.Atoi(rangeSpanStr[1])

		for i := rangeSpan[0]; i <= rangeSpan[1]; i++ {
			numberStr := strconv.Itoa(i)
			digits := len(numberStr)

			if digits%2 == 1 {
				continue //Skipping odd numbers as they cannot have "twice" patterns
			}

			middle := digits / 2

			invalidId := true

			for k := 0; k < middle; k++ {
				if numberStr[k] != numberStr[middle+k] {
					invalidId = false
					break
				}
			}

			if invalidId {
				fmt.Println(numberStr, "-", digits)
				sum += i
			}

		}
	}

	return sum
}

func solvePart2(lines []string) int {
	return 0
}
