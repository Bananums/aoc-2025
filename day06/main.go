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

	lines, err := util.LoadFile("puzzle.txt", inputs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", solvePart1(lines, false))
	fmt.Println("Part 2:", solvePart2(lines, false))
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

func solvePart2(lines []string, verbose bool) int {
	sum := 0

	height := len(lines)
	width := 0
	for _, line := range lines {
		if len(line) > width {
			width = len(line)
		}
	}

	if verbose {
		for _, line := range lines {
			fmt.Println(line)
		}
		fmt.Println("------------------------")
	}

	splits := make([][]byte, height)
	for i := range splits {
		splits[i] = make([]byte, width)
	}

	for i := 0; i < height; i++ {
		currentWidth := len(lines[i])
		for j := 0; j < currentWidth; j++ {
			if lines[i][j] == ' ' {
				splits[i][j] = '.'
			} else {
				splits[i][j] = lines[i][j]
			}
		}
		for currentWidth < width {
			splits[i][currentWidth] = '.'
			currentWidth++
		}
	}

	//Just printing for check
	if verbose {
		for _, line := range splits {
			fmt.Println(string(line))
		}
	}

	var operator byte
	sumPart := 0
	for i := 0; i < width; i++ {
		if splits[height-1][i] == '+' || splits[height-1][i] == '*' {
			operator = splits[height-1][i]
		}

		numberStr := ""
		for k := 0; k < height-1; k++ { // -1 to skip operator * and +
			numberStr += string(splits[k][i])
		}
		numberStr = strings.Trim(numberStr, ".")
		number, _ := strconv.Atoi(numberStr)

		if verbose {
			fmt.Println(number, "len:", len(numberStr), "operator:", string(operator))
		}

		if len(numberStr) == 0 {
			if verbose {
				fmt.Println("sumPart: ", sumPart)
			}
			sum += sumPart
			sumPart = 0
			continue
		}

		if operator == '+' {
			sumPart += number
		} else {
			if sumPart == 0 {
				sumPart = 1
			}
			sumPart *= number
		}
	}

	if verbose {
		fmt.Println("sumPart: ", sumPart)
	}
	sum += sumPart // I do not care anymore. It works

	return sum
}

type Column struct {
	number int
	op     byte
}
