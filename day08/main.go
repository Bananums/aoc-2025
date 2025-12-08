package main

import (
	"aoc-2025/internal/util"
	"embed"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed example.txt puzzle.txt
var inputs embed.FS

func main() {
	fmt.Println("Advent of Code - Day 08")

	lines, err := util.LoadFile("example.txt", inputs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", solvePart1(lines, false))
	fmt.Println("Part 2:", solvePart2(lines, false))
}

func solvePart1(lines []string, verbose bool) int {
	_ = verbose

	junctions := make([]Junction, len(lines))
	for i, line := range lines {
		junctions[i] = makeJunction(line)
	}

	used := make(map[[2]int]bool)
	linkChecks := 10
	nextGridId := 1

	_ = used

	for check := 0; check < linkChecks; check++ {
		minDistance := math.MaxFloat64
		currentMinLink := [2]int{-1, -1}
		for i := range junctions {
			for j := range junctions {
				if j == i {
					continue //No point in checking self
				}
				k := getKey(i, j)
				if used[k] {
					fmt.Println("Skipping", k, "-", junctions[i], junctions[j])
				}

				distance := getDistance(junctions[j], junctions[i])
				if distance < minDistance {
					minDistance = distance
					currentMinLink = [2]int{i, j}
				}
			}
		}
		junctions[currentMinLink[0]].grid = nextGridId
		junctions[currentMinLink[1]].grid = nextGridId
		used[currentMinLink] = true
		fmt.Println("Min:", minDistance, junctions[currentMinLink[0]], junctions[currentMinLink[1]])
	}

	return 0
}

func solvePart2(lines []string, verbose bool) int {
	_ = verbose
	return 0
}

type Junction struct {
	x    float64
	y    float64
	z    float64
	grid int
}

func getKey(i int, j int) [2]int {
	//Makes sure pairs are made symmetrical
	// (i, j) and (j, i) should count as the same pair
	if i < j {
		return [2]int{i, j}
	}
	return [2]int{j, i}
}

func makeJunction(line string) Junction {
	trimmedLine := strings.Split(line, ",")
	numberX, _ := strconv.Atoi(trimmedLine[0])
	numberY, _ := strconv.Atoi(trimmedLine[1])
	numberZ, _ := strconv.Atoi(trimmedLine[2])
	return Junction{float64(numberX), float64(numberY),
		float64(numberZ), 0}
}

func getDistance(start Junction, end Junction) float64 {
	dx := end.x - start.x
	dy := end.y - start.y
	dz := end.z - start.z
	d := math.Sqrt(dx*dx + dy*dy + dz*dz)

	//It is okay to compare squared distances because the square root
	//function is monotonically increasing for non-negative numbers.
	//Thank you for coming to my ted talk.
	//TODO Validate when solution

	return d
}
