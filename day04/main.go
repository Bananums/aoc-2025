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
	fmt.Println("Advent of Code - Day 04")

	lines, err := util.LoadFile("puzzle.txt", inputs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", solvePart1(lines, false))
	fmt.Println("Part 2:", solvePart2(lines))
}

func solvePart1(lines []string, verbose bool) int {

	// Generate an indexable matrix Start ------------------------------------------------------------
	a := len(lines)    //Matrix height
	b := len(lines[0]) //Matrix width

	validA := [2000]int{}
	for i := range validA {
		validA[i] = -1
	}
	validB := [2000]int{}
	for i := range validB {
		validB[i] = -1
	}
	validPosition := 0

	data := make([]byte, a*b)
	matrix := make([][]byte, a)
	for i := 0; i < a; i++ {
		matrix[i] = data[i*b : (i+1)*b]
	}

	for i, line := range lines {
		for k, char := range line {
			matrix[i][k] = byte(char)
		}
	}

	if verbose {
		for i := 0; i < a; i++ {
			fmt.Println(string(matrix[i]))
		}
	}
	// Generate an indexable matrix End ------------------------------------------------------------

	validScrolls := 0
	for i := 0; i < a; i++ {
		for k := 0; k < b; k++ {
			scrollCount := -1
			if matrix[i][k] == '@' {
				scrollCount = 0
				if i != 0 {
					if matrix[i-1][k] == '@' {
						scrollCount++
					} // Above

					if k > 0 {
						if matrix[i-1][k-1] == '@' {
							scrollCount++
						} // Above Left
					}

					if k < b-1 {
						if matrix[i-1][k+1] == '@' {
							scrollCount++
						} // Above Right
					}
				}

				if k > 0 {
					if matrix[i][k-1] == '@' {
						scrollCount++
					} // Left
				}

				if k < b-1 {
					if matrix[i][k+1] == '@' {
						scrollCount++
					} // Right
				}

				if i < a-1 {
					if matrix[i+1][k] == '@' {
						scrollCount++
					} // Below

					if k > 0 {
						if matrix[i+1][k-1] == '@' {
							scrollCount++
						} // Below Left
					}

					if k < b-1 {
						if matrix[i+1][k+1] == '@' {
							scrollCount++
						} // Below Right
					}
				}

			}
			if scrollCount < 4 && scrollCount != -1 {
				validA[validPosition] = i
				validB[validPosition] = k
				validPosition++
				validScrolls++
			}
		}

	}

	for i := range validPosition {
		w := validA[i]
		l := validB[i]
		if w != -1 && l != -1 {
			matrix[w][l] = 'x'
		}
	}

	if verbose {
		fmt.Println("-----------")
		for i := 0; i < a; i++ {
			fmt.Println(string(matrix[i]))
		}
	}

	return validScrolls
}

func solvePart2(lines []string) int {
	matrix := GenerateBaseMatrix(lines)
	totalScrolls := 0
	scrolls := 1
	for scrolls > 0 {
		matrix, scrolls = RemoveScrolls(matrix)
		totalScrolls += scrolls
	}

	return totalScrolls
}

func GenerateBaseMatrix(lines []string) [][]byte {
	a := len(lines)    //Matrix height
	b := len(lines[0]) //Matrix width

	data := make([]byte, a*b)
	matrix := make([][]byte, a)
	for i := 0; i < a; i++ {
		matrix[i] = data[i*b : (i+1)*b]
	}

	for i, line := range lines {
		for k, char := range line {
			matrix[i][k] = byte(char)
		}
	}

	return matrix
}

func RemoveScrolls(matrix [][]byte) ([][]byte, int) {

	a := len(matrix)
	b := len(matrix[0])

	validA := [2000]int{}
	for i := range validA {
		validA[i] = -1
	}
	validB := [2000]int{}
	for i := range validB {
		validB[i] = -1
	}
	validPosition := 0

	validScrolls := 0
	for i := 0; i < a; i++ {
		for k := 0; k < b; k++ {
			scrollCount := -1
			if matrix[i][k] == '@' {
				scrollCount = 0
				if i != 0 {
					if matrix[i-1][k] == '@' {
						scrollCount++
					} // Above

					if k > 0 {
						if matrix[i-1][k-1] == '@' {
							scrollCount++
						} // Above Left
					}

					if k < b-1 {
						if matrix[i-1][k+1] == '@' {
							scrollCount++
						} // Above Right
					}
				}

				if k > 0 {
					if matrix[i][k-1] == '@' {
						scrollCount++
					} // Left
				}

				if k < b-1 {
					if matrix[i][k+1] == '@' {
						scrollCount++
					} // Right
				}

				if i < a-1 {
					if matrix[i+1][k] == '@' {
						scrollCount++
					} // Below

					if k > 0 {
						if matrix[i+1][k-1] == '@' {
							scrollCount++
						} // Below Left
					}

					if k < b-1 {
						if matrix[i+1][k+1] == '@' {
							scrollCount++
						} // Below Right
					}
				}

			}
			if scrollCount < 4 && scrollCount != -1 {
				validA[validPosition] = i
				validB[validPosition] = k
				validPosition++
				validScrolls++
			}
		}

	}

	for i := range validPosition {
		w := validA[i]
		l := validB[i]
		if w != -1 && l != -1 {
			matrix[w][l] = '.'
		}
	}

	return matrix, validScrolls
}
