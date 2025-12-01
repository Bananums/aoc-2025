package main

import (
	"aoc-2025/internal/util"
	"embed"
	"fmt"
	"log"
)

//go:embed example.txt
var inputs embed.FS

func main() {
	fmt.Println("Advent of Code - Day XX")

	lines, err := util.LoadFile("example.txt", inputs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", solvePart1(lines))
	fmt.Println("Part 2:", solvePart2(lines))
}

func solvePart1(lines []string) int {
	return 0
}

func solvePart2(lines []string) int {
	return 0
}
