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
	fmt.Println("Advent of Code - Day 06")

	lines, err := util.LoadFile("example.txt", inputs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", solvePart1(lines, false))
	fmt.Println("Part 2:", solvePart2(lines))
}

func solvePart1(lines []string, verbose bool) int {
	splits := make([][]string, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)

		if verbose {
			fmt.Println(fields)
		}

		splits[i] = fields
	}

	a := len(splits)
	b := len(splits[0])

	transposed := make([][]string, b)
	for i := range b {
		transposed[i] = make([]string, a)
	}

	shift := 0
	for shift < b {
		for i := range a {
			transposed[shift][i] = splits[i][shift]
		}
		shift++
	}

	sum := 0
	for _, fields := range transposed {
		calculationLength := len(fields) - 1
		taskValue := 0
		for i := 0; i < calculationLength; i++ {
			value, _ := strconv.Atoi(fields[i])
			if fields[len(fields)-1] == "*" {
				if taskValue == 0 {
					taskValue = 1
				}
				taskValue *= value
			} else {
				taskValue += value
			}
		}
		sum += taskValue
	}

	return sum
}

func solvePart2(lines []string) int {

	splits := make([][]string, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)
		splits[i] = fields
	}

	for _, split := range splits {
		fmt.Println(split)
	}

	width := len(splits[0])
	height := len(splits)
	//fmt.Println("width:", width, "height:", height)

	rows := make([][]string, width)
	for i := range rows {
		rows[i] = make([]string, height)
	}

	//fmt.Println("rows height:", width, "rows width:", height)

	for i := 0; i < height; i++ {
		for k := 0; k < width; k++ {
			rows[k][i] = splits[i][k]
		}
	}

	for _, row := range rows {
		fmt.Println("--------------")
		fmt.Println(row)
		largestSize := 0
		for _, fields := range row {
			if len(fields) > largestSize {
				largestSize = len(fields)
			}
		}

		columns := make([][]byte, largestSize)
		for i := range columns {
			columns[i] = make([]byte, len(rows)-1) // Removing operator * or +
		}

		for i, column := range columns {
			for j := range column {
				numberFromRow := row[j] // Getting number from row[j]
				numberLength := len(numberFromRow)
				index := numberLength - 1 - i // Counting right to left

				if index >= 0 {
					columns[i][j] = numberFromRow[index]
				} else {
					columns[i][j] = '.'
				}
			}
		}

		fmt.Println("printing columns")
		//valueStr := []string{} // TODO Make string and then convert to int
		for _, column := range columns {
			fmt.Println(string(column))

			for _, value := range column {
				fmt.Println(string(value))
			}

		}

	}

	return 0
}
