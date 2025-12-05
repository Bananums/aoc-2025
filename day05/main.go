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
	fmt.Println("Advent of Code - Day 05")

	idRanges, ingredients, err := util.LoadSplitFile("puzzle.txt", inputs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", solvePart1(idRanges, ingredients, false))
	fmt.Println("Part 2:", solvePart2(idRanges, ingredients))
}

func solvePart1(idRanges []string, ingredients []string, verbose bool) int {
	freshIngredients := 0

	for _, ingredient := range ingredients {

		ingredientAsInt, _ := strconv.Atoi(ingredient)

		for _, idRange := range idRanges {
			idStart, idEnd, err := parseRange(idRange)
			if err != nil {
			}

			if ingredientAsInt < idStart {
				if verbose {
					fmt.Println("[", idStart, "-", idEnd, "] - ", ingredientAsInt, " - spoiled")
				}
				continue
			}

			if ingredientAsInt > idEnd {
				if verbose {
					fmt.Println("[", idStart, "-", idEnd, "] - ", ingredientAsInt, " - spoiled")
				}
				continue
			}

			if verbose {
				fmt.Println("[", idStart, "-", idEnd, "] - ", ingredientAsInt, " - fresh")
			}

			freshIngredients++
			break
		}

	}

	return freshIngredients
}

func solvePart2(idRanges []string, ingredients []string) int {
	// Read this. Figure out how to implement Hasmap. Then blast through with lookup
	// https://medium.com/@hillman.t.y/designing-a-hashmap-from-scratch-51870851406f
	return 0
}

func parseRange(idRange string) (int, int, error) {
	parts := strings.Split(idRange, "-")

	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid range: %q", idRange)
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}

	return start, end, nil
}
