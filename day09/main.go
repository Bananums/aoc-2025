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

	lines, err := util.LoadFile("example.txt", inputs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", solvePart1(lines, false))
	fmt.Println("Part 2:", solvePart2(lines, true))
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
	points := linesAsPoints(lines)

	if verbose {
		fmt.Println(points)
		fmt.Println("-------------------------------")
	}

	size := 0
	largestPoint1 := Point{X: 0, Y: 0}
	largestPoint2 := Point{X: 0, Y: 0}
	_ = largestPoint1
	_ = largestPoint2

	for i := range points {
		for j := range points {
			point1 := points[i]
			point2 := points[j]

			if point1.X == point2.X && point1.Y == point2.Y {
				continue
			}

			invCorner1 := Point{points[i].Y, points[j].Y}
			invCorner2 := Point{points[j].X, points[i].X}

			//fmt.Println(point1, point2, invCorner1, invCorner2)

			for _, point := range points {
				if point == invCorner1 || point == invCorner2 {
					area := (point2.X - point1.X + 1) * (point2.Y - point1.Y + 1)
					if area > size {
						size = area
						fmt.Println(point1, point2, invCorner1, invCorner2)
						largestPoint1 = point1
						largestPoint2 = point2
					}
				}
			}
		}
	}

	var grid [9][14]byte
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for _, point := range points {
		grid[point.Y][point.X] = '#' // Y then X since program goes height then width in direction
	}

	//for i := largestPoint1.X; i < largestPoint2.X+1; i++ {
	//	for j := largestPoint1.Y; j < largestPoint2.Y+1; j++ {
	//		grid[j][i] = 'O'
	//	}
	//}

	fmt.Println("------------------------------")

	for i := range points {
		for j := range points {
			if points[i] == points[j] {
				continue
			}

			if points[j].X-points[i].X < 0 {
				continue
			}

			if points[i].Y == points[j].Y {
				fmt.Println(points[i], points[j])
				for k := points[i].X + 1; k < points[j].X; k++ {
					grid[j][k] = 'X'
				}
			}
		}
	}

	for _, row := range grid {
		fmt.Println(string(row[:]))
	}

	return size
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
