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
	fmt.Println("Advent of Code - Day 07")

	lines, err := util.LoadFile("puzzle.txt", inputs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", solvePart1(lines, false))
	fmt.Println("Part 2:", solvePart2(lines, false))
}

func solvePart1(lines []string, verbose bool) int {
	_ = verbose

	myLines := toByteLines(lines)
	for _, line := range myLines {
		fmt.Println(string(line))
	}

	fmt.Println("+---------------+")

	startIndex := findStartIndex(myLines[0])
	fmt.Println("Start index:", startIndex)

	fmt.Println(string(myLines[0]))
	splits := 0
	for i := 1; i < len(myLines); i++ {
		if i == 1 {
			myLines[i][startIndex] = '|'
			fmt.Println(string(myLines[i]))
			continue
		}
		if i%2 == 0 { // Even number
			for k, char := range string(myLines[i]) {
				if myLines[i-1][k] == '|' {
					if char == '^' {
						myLines[i][k-1] = '|'
						myLines[i][k+1] = '|'
						splits++
					} else {
						myLines[i][k] = '|'
					}
				}
			}
		} else {
			for k := range string(myLines[i]) {
				if myLines[i-1][k] == '|' {
					myLines[i][k] = '|'
				}
			}
		}

		fmt.Println(string(myLines[i]))
	}

	return splits
}

func solvePart2(lines []string, verbose bool) int {
	return 0
}

func toByteLines(lines []string) [][]byte {
	out := make([][]byte, len(lines))
	for i, s := range lines {
		out[i] = []byte(s)
	}
	return out
}

func findStartIndex(line []byte) int {
	startIndex := -1
	startChar := byte('S')
	for index, char := range line {
		if char == startChar {
			startIndex = index
			break
		}
	}
	return startIndex
}
