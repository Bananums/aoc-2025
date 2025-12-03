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

	fmt.Println("Part 1:", solvePart1(lines))
	fmt.Println("Part 2:", solvePart2(lines))
}

func solvePart1(lines []string) int {

	joltage := 0

	for _, line := range lines {
		//fmt.Println(line)

		var largest uint8 = '0'
		var largestIndex = 0

		var largestSecond uint8 = '0'
		//var largestIndexSecond = 0

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
				//largestIndexSecond = index + 1
			}
		}

		//fmt.Println(string(largest), "-", largestIndex, "-", string(largestSecond), "-", largestIndexSecond)

		value := (largest-'0')*10 + (largestSecond - '0')
		//fmt.Println("Value:", value)
		joltage += int(value)
	}

	return joltage
}

func solvePart2(lines []string) int {
	return 0
}
