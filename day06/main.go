package main

import (
	"aoc-2025/internal/util"
	"embed"
	"fmt"
	"log"
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

	fmt.Println("Part 1:", solvePart1(lines))
	fmt.Println("Part 2:", solvePart2(lines))
}

func solvePart1(lines []string) int {

	test1 := string("bob")

	fmt.Println(test1)
	fmt.Println("----------------")

	test2 := make([]string, 3)
	test2[0] = "bob"
	test2[1] = "alex"
	test2[2] = "john"

	for _, yep := range test2 {
		fmt.Println(yep)
	}
	fmt.Println("----------------")

	var test3 [3][3]string

	//test3 := make([][]string, 3)
	//for i := range test3 {
	//	test3[i] = make([]string, 3)
	//}

	test3[0][0] = "bob1"
	test3[0][1] = "bob2"
	test3[0][2] = "bob3"
	test3[1][0] = "alex1"
	test3[1][1] = "alex2"
	test3[1][2] = "alex3"
	test3[2][0] = "john1"
	test3[2][1] = "john2"
	test3[2][2] = "john3"

	for _, yep := range test3 {
		fmt.Println(yep)
	}

	fmt.Println("----------------")

	splits := make([][]string, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)
		fmt.Println(fields)
		splits[i] = fields
	}

	for _, fields := range splits {
		fmt.Println(fields)
	}

	return 0
}

func solvePart2(lines []string) int {
	return 0
}
