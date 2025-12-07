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

	lines, err := util.LoadFile("example.txt", inputs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", solvePart1(lines, false))
	fmt.Println("Part 2:", solvePart2(lines, true))
}

func solvePart1(lines []string, verbose bool) int {

	myLines := toByteLines(lines)
	startIndex := findStartIndex(myLines[0])

	if verbose {
		for _, line := range myLines {
			fmt.Println(string(line))
		}
		fmt.Println("+---------------+")
		fmt.Println(string(myLines[0]))
	}

	splits := 0
	for i := 1; i < len(myLines); i++ {
		if i == 1 {
			myLines[i][startIndex] = '|'
			if verbose {
				fmt.Println(string(myLines[i]))
			}
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

		if verbose {
			fmt.Println(string(myLines[i]))
		}
	}

	return splits
}

func solvePart2(lines []string, verbose bool) int {
	myLines := toByteLines(lines)
	pascalGrid := makeIntGrid(lines)
	startIndex := findStartIndex(myLines[0])

	if verbose {
		for _, line := range myLines {
			fmt.Println(string(line))
		}
		fmt.Println("+---------------+")
		for _, line := range pascalGrid {
			fmt.Println(line)
		}
		fmt.Println("+---------------+")

		fmt.Println(string(myLines[0]))
	}

	for i := 1; i < len(myLines); i++ {
		if i == 1 {
			myLines[i][startIndex] = '|'
			if verbose {
				fmt.Println(string(myLines[i]))
			}
			continue
		}
		if i%2 == 0 { // Even number
			for k, char := range string(myLines[i]) {
				if myLines[i-1][k] == '|' {
					if char == '^' {
						myLines[i][k-1] = '|'
						myLines[i][k+1] = '|'

						matches := 0
						if i == 2 {
							matches = 1
						}
						if myLines[i-2][k-1] == '^' {
							matches += pascalGrid[i-2][k-1]
						}
						if myLines[i-2][k+1] == '^' {
							matches += pascalGrid[i-2][k+1]
						}
						pascalGrid[i][k] = matches
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

		if verbose {
			fmt.Println(string(myLines[i]))
		}
	}

	if verbose {
		for i, line := range pascalGrid {
			if i == 0 {
				continue
			}
			if i%2 == 0 {
				fmt.Println(line)
			}
		}
	}

	splits := 0
	for _, line := range pascalGrid {
		for _, value := range line {
			splits += value
		}
	}

	return splits
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

func makeIntGrid(lines []string) [][]int {
	out := make([][]int, len(lines))
	for i := range lines {
		out[i] = make([]int, len(lines[i]))
	}
	return out
}
