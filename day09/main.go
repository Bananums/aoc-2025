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
	fmt.Println("Advent of Code - Day 09")

	lines, err := util.LoadFile("puzzle.txt", inputs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", solvePart1(lines, false))
	fmt.Println("Part 2:", solvePart2(lines, false))
}

func solvePart1(lines []string, verbose bool) int {

	points := linesAsPoints(lines)

	if verbose {
		fmt.Println(points)
	}

	size := 0
	for i := range points {
		for j := range points {
			area := (points[j].X - points[i].X + 1) * (points[j].Y - points[i].Y + 1)
			if area > size {
				size = area
			}
		}
	}

	return size
}

func solvePart2(lines []string, verbose bool) int {
	_ = verbose
	return 0
}

func linesAsPoints(lines []string) []Point {
	points := make([]Point, len(lines))
	for i, line := range lines {
		split := strings.Split(line, ",")
		if len(split) != 2 {
			log.Fatal("Line ", line, " doesn't contain 2 points")
		}

		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])

		point := Point{x, y}
		points[i] = point
	}
	return points
}

type Point struct {
	X int
	Y int
}
