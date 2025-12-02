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

	fmt.Println("Part 1:", solvePart1(lines, false))
	fmt.Println("Part 2:", solvePart2(lines, false))
}

func solvePart1(lines []string, verbose bool) int {

	sum := 0

	for _, line := range lines {
		if verbose {
			fmt.Println(line)
		}

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
				if verbose {
					fmt.Println(numberStr, "-", digits)
				}
				sum += i
			}

		}
	}

	return sum
}

func solvePart2(lines []string, verbose bool) int {

	sum := 0

	for _, line := range lines {
		rangeSpanStr := strings.Split(line, "-")
		var rangeSpan [2]int
		rangeSpan[0], _ = strconv.Atoi(rangeSpanStr[0])
		rangeSpan[1], _ = strconv.Atoi(rangeSpanStr[1])

		for i := rangeSpan[0]; i <= rangeSpan[1]; i++ {
			numberStr := strconv.Itoa(i)
			digits := len(numberStr)
			sum += BruteForceSearch(i, numberStr, digits, verbose)
		}
	}

	return sum
}

func BruteForceSearch(number int, numberStr string, digits int, verbose bool) int {
	sum := 0
	approvedPattern := ""

	//Only checking half into the number, since a sequence cannot exist with more than that
	for i := 1; i <= digits/2; i++ {
		if numberStr == approvedPattern {
			break
		}

		if digits%i != 0 {
			continue
		}

		var ok = true
		for k := 0; k < digits; k++ {
			if numberStr[k] != numberStr[k%i] {
				ok = false
				break
			}
		}

		if !ok {
			continue
		}

		reappearance := digits / i
		if reappearance < 2 {
			continue
		}

		if verbose {
			fmt.Println(numberStr)
		}
		sum += number
		approvedPattern = numberStr
	}

	return sum
}
