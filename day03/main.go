package main

import (
	"aoc-2025/internal/util"
	"embed"
	"fmt"
	"log"
)

//go:embed example.txt puzzle.txt
var inputs embed.FS

func main() {
	fmt.Println("Advent of Code - Day 03")

	lines, err := util.LoadFile("puzzle.txt", inputs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", solvePart1(lines, false))
	fmt.Println("Part 2:", solvePart2(lines, false))
}

func solvePart1(lines []string, verbose bool) int {

	joltage := 0

	for _, line := range lines {
		//fmt.Println(line)

		var largest uint8 = '0'
		var largestIndex = 0

		var largestSecond uint8 = '0'
		var largestIndexSecond = 0

		for index := range line {
			if line[index] > largest && index < len(line)-1 {
				largest = line[index]
				largestIndex = index
			}
		}

		sub := line[largestIndex+1:]
		for index := range sub {
			if sub[index] > largestSecond {
				largestSecond = sub[index]
				largestIndexSecond = index + 1
			}
		}

		if verbose {
			fmt.Println(string(largest), "-", largestIndex, "-",
				string(largestSecond), "-", largestIndexSecond)
		}

		value := (largest-'0')*10 + (largestSecond - '0')
		if verbose {
			fmt.Println("Value:", value)
		}
		joltage += int(value)
	}

	return joltage
}

func solvePart2(lines []string, verbose bool) int64 {

	var joltage int64 = 0
	const digitSize = 12

	for _, line := range lines {
		allowedRemovals := len(line) - digitSize

		//Just allocating 100 int32 numbers
		var result [100]int32
		var resultHuman [100]int32

		position := 0
		for i, digit := range line {

			//Add the first number to get started
			if i == 0 {
				result[0] = digit
				resultHuman[0] = digit - '0'
				continue
			}

			for position >= 0 && allowedRemovals > 0 && result[position] < digit {
				position--
				allowedRemovals--
			}

			position++
			result[position] = digit
			resultHuman[position] = digit - '0'

		}

		var movedResult [digitSize]int32
		for i, digit := range result[:digitSize] {
			movedResult[i] = digit
		}

		if verbose {
			fmt.Println(line, "-->", resultHuman, "-", twelveAsciiDigitsToNumber(movedResult))
		}

		joltage += twelveAsciiDigitsToNumber(movedResult)
	}

	return joltage
}

func twelveAsciiDigitsToNumber(digits [12]int32) int64 {
	var number int64
	for _, digit := range digits {
		number = number*10 + int64(digit-'0')
	}
	return number
}
