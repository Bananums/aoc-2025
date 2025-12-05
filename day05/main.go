package main

import (
	"aoc-2025/internal/util"
	"embed"
	"fmt"
	"log"
	"sort"
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
	fmt.Println("Part 2:", solvePart2(idRanges, ingredients, false))
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

type Interval struct {
	Start uint64
	End   uint64
}

func solvePart2(idRanges []string, ingredients []string, verbose bool) int {
	_ = verbose
	freshIngredients := 0
	var intervals = GetIntervals(idRanges)

	//Sort the list to make everything easy peasy lemon squeezy.
	sort.Slice(intervals, func(i, j int) bool { return intervals[i].Start < intervals[j].Start })

	merged := MergeIntervals(intervals)
	for _, interval := range merged {
		diff := interval.End - interval.Start + 1 // does not work if range is e.g. 5-5
		freshIngredients += int(diff)
	}

	return freshIngredients
}

func MergeIntervals(intervals []Interval) []Interval {
	var merged []Interval
	for i, current := range intervals {

		if i == 0 {
			merged = append(merged, current)
			continue
		}

		last := &merged[len(merged)-1]

		if current.Start <= last.End {
			// overlap -> extend
			if current.End > last.End {
				last.End = current.End
			}
		} else {
			// no overlap -> new interal
			merged = append(merged, current)
		}

	}
	return merged
}

func GetIntervals(idRanges []string) []Interval {
	var intervals []Interval
	for _, idRange := range idRanges {
		idStart, idEnd, _ := parseRangeUint64(idRange)
		intervals = append(intervals, Interval{idStart, idEnd})
	}
	return intervals
}

func parseRangeUint64(idRange string) (uint64, uint64, error) {
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

	return uint64(start), uint64(end), nil
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
