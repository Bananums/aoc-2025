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

	filename := "puzzle.txt"

	lines, err := util.LoadFile(filename, inputs)
	if err != nil {
		log.Fatal(err)
	}

	if filename == "example.txt" {
		fmt.Println("Part 1:", solvePart1(lines, false, 10))
	} else {
		fmt.Println("Part 1:", solvePart1(lines, false, 1000))
	}

	fmt.Println("Part 2:", solvePart2(lines, false))

}

func solvePart1(lines []string, verbose bool, checks int) int {
	_ = verbose

	junctions := make([]Junction, len(lines))
	for i, line := range lines {
		junctions[i] = makeJunction(line)
	}

	used := make(map[[2]int]bool)
	linkChecks := checks
	nextGridId := 1

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
					//fmt.Println("Skipping", k, "-", junctions[i], junctions[j])
					continue
				}

				distance := getDistance(junctions[j], junctions[i])
				if distance < minDistance {
					minDistance = distance
					currentMinLink = [2]int{i, j}
				}
			}
		}

		nearestIndexI := currentMinLink[0]
		nearestIndexJ := currentMinLink[1]

		if nearestIndexI == -1 {
			// no valid edge left
			break
		}
		used[getKey(nearestIndexI, nearestIndexJ)] = true

		gi := junctions[nearestIndexI].grid
		gj := junctions[nearestIndexJ].grid

		switch {
		case gi == 0 && gj == 0:
			// New circuit with two junctions
			junctions[nearestIndexI].grid = nextGridId
			junctions[nearestIndexJ].grid = nextGridId
			nextGridId++

		case gi != 0 && gj == 0:
			// Add J to I's circuit
			junctions[nearestIndexJ].grid = gi

		case gi == 0 && gj != 0:
			// Add I to J's circuit
			junctions[nearestIndexI].grid = gj

		case gi != 0 && gj != 0 && gi != gj:
			// MERGE two circuits
			oldId := gj
			newId := gi
			for idx := range junctions {
				if junctions[idx].grid == oldId {
					junctions[idx].grid = newId
				}
			}

		case gi == gj:
			// Already in the same circuit: do nothing to grids
			// but still counts as one of the 10
		}

		//fmt.Println("Min:", minDistance, junctions[currentMinLink[0]], junctions[currentMinLink[1]])
	}

	var grids [1000]int

	for _, junction := range junctions {
		//fmt.Println(junction)
		grids[junction.grid] += 1
	}

	//fmt.Println(grids)

	val1 := 0
	val2 := 0
	val3 := 0

	for i := 1; i < len(grids); i++ {
		value := grids[i]

		if value > val1 {
			val3 = val2
			val2 = val1
			val1 = value
		} else if value > val2 {
			val3 = val2
			val2 = grids[i]
		} else if value > val3 {
			val3 = value
		}

	}

	//fmt.Println(val1, val2, val3)

	return val1 * val2 * val3
}

func solvePart2(lines []string, verbose bool) int {
	//Can use union–find (DSU) for a cleaner implementation. Look into this
	_ = verbose

	junctions := make([]Junction, len(lines))
	for i, line := range lines {
		junctions[i] = makeJunction(line)
	}

	used := make(map[[2]int]bool)
	nextGridId := 1

	circuits := len(junctions)
	lastMerge := [2]int{-1, -1}

	for check := 0; circuits > 1; check++ {
		minDistance := math.MaxFloat64
		currentMinLink := [2]int{-1, -1}
		for i := range junctions {
			for j := range junctions {
				if j == i {
					continue //No point in checking self
				}

				k := getKey(i, j)
				if used[k] {
					//fmt.Println("Skipping", k, "-", junctions[i], junctions[j])
					continue
				}

				distance := getDistance(junctions[j], junctions[i])
				//fmt.Println(distance, " - ", junctions[i], junctions[j])
				if distance < minDistance {
					minDistance = distance
					currentMinLink = [2]int{i, j}
				}
			}
		}

		nearestIndexI := currentMinLink[0]
		nearestIndexJ := currentMinLink[1]

		if nearestIndexI == -1 {
			// no valid edge left
			break
		}
		used[getKey(nearestIndexI, nearestIndexJ)] = true

		gi := junctions[nearestIndexI].grid
		gj := junctions[nearestIndexJ].grid

		switch {
		case gi == 0 && gj == 0:
			// New circuit with two junctions
			junctions[nearestIndexI].grid = nextGridId
			junctions[nearestIndexJ].grid = nextGridId
			nextGridId++
			circuits--
			lastMerge = currentMinLink

		case gi != 0 && gj == 0:
			// Add J to I's circuit
			junctions[nearestIndexJ].grid = gi
			circuits--
			lastMerge = currentMinLink

		case gi == 0 && gj != 0:
			// Add I to J's circuit
			junctions[nearestIndexI].grid = gj
			circuits--
			lastMerge = currentMinLink

		case gi != 0 && gj != 0 && gi != gj:
			// MERGE two circuits
			//fmt.Println("Mering:", junctions[gi], junctions[gj])
			oldId := gj
			newId := gi
			for idx := range junctions {
				if junctions[idx].grid == oldId {
					junctions[idx].grid = newId
				}
			}
			circuits--
			lastMerge = currentMinLink

		case gi == gj:
			// Already in the same circuit: do nothing to grids
			// but still counts as one of the 10
		}

		//fmt.Println("Min:", minDistance, junctions[currentMinLink[0]], junctions[currentMinLink[1]])
	}

	if lastMerge[0] != -1 {
		fmt.Println("Last merge that connects all junctions:",
			junctions[lastMerge[0]], junctions[lastMerge[1]])
	}

	return int(junctions[lastMerge[0]].x * junctions[lastMerge[1]].x)
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
