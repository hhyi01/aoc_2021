package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func day07Part1(filePath string)  {
	puzzleInput := readFileLines(filePath)
	positions := strings.Split(puzzleInput[0], ",")
	minPos := calcMinPosition(positions)
	maxPos := calcMaxPosition(positions)
	fuelCalcs := calcFuelOnPosition(positions, minPos, maxPos, "Part 1")
	leastFuelPos, fuelConsumed := calcMinFuel(fuelCalcs)
	fmt.Println("Horizontal position using least fuel =", leastFuelPos)
	fmt.Println("Part 1 - fuel spent to align to horizontal position =", fuelConsumed)
}

func day07Part2(filePath string)  {
	puzzleInput := readFileLines(filePath)
	positions := strings.Split(puzzleInput[0], ",")
	minPos := calcMinPosition(positions)
	maxPos := calcMaxPosition(positions)
	fuelCalcs := calcFuelOnPosition(positions, minPos, maxPos, "Part 2")
	leastFuelPos, fuelConsumed := calcMinFuel(fuelCalcs)
	fmt.Println("Horizontal position using least fuel =", leastFuelPos)
	fmt.Println("Part 2 - fuel spent to align to horizontal position =", fuelConsumed)
}

func calcFuelOnPosition(positions []string, minPos int, maxPos int, calcType string) map[int]int {
	fuelCalcs := make(map[int]int)
	for pos := minPos; pos <= maxPos; pos++ {
		totalFuel := 0
		for _, v := range positions {
			crabSubPos, _ := strconv.Atoi(v)
			_, found := fuelCalcs[pos]
			if !found {
				fuel := pos - crabSubPos
				if fuel < 0 {
					fuel = fuel * -1
				}
				if calcType == "Part 1" {
					totalFuel += fuel
				} else {
					crabFuelCost := crabCalcFuel(1, fuel, fuel)
					totalFuel += crabFuelCost
				}
			}
		}
		_, found := fuelCalcs[pos]
		if !found {
			fuelCalcs[pos] = totalFuel
		}
	}
	return fuelCalcs
}

func calcMinPosition(crabPositions []string) int {
	minPos := math.MaxInt
	for _, v := range crabPositions {
		pos, _ := strconv.Atoi(v)
		if pos < minPos {
			minPos = pos
		}
	}
	return minPos
}

func calcMaxPosition(crabPositions []string) int {
	maxPos := math.MinInt
	for _, v := range crabPositions {
		pos, _ := strconv.Atoi(v)
		if pos > maxPos {
			maxPos = pos
		}
	}
	return maxPos
}

func calcMinFuel(fuelCalcs map[int]int) (int, int) {
	minFuel := math.MaxInt
	minFuelPos := 0
	for k, v := range fuelCalcs {
		if v < minFuel {
			minFuel = v
			minFuelPos = k
		}
	}
	return minFuelPos, minFuel
}

func crabCalcFuel(start int, end int, numTerms int) int {
	sum := numTerms * (start + end) / 2
	return sum
}