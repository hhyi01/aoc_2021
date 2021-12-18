package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func day17Part1And2(filePath string) {
	puzzleInput := readFileLines(filePath)
	fmt.Println(puzzleInput)
	targetArea := setTargetArea(puzzleInput)
	maxY := math.MinInt
	maxInitVelocity := []int{0,0}
	iterations := 0
	uniqueHits := make(map[string]bool)
	for x := 1; x <= targetArea["x"][1]; x++ {
		for y := targetArea["y"][0]*2; y <= (targetArea["y"][0] * -1)*2; y++ {
			initVelocity := []int{x,y}
			initPosition := initVelocity
			localMaxY := math.MinInt
			velocity := initVelocity
			position := initPosition
			for {
				onTarget := inTargetArea(position, targetArea)
				nextVelocity := calculateVelocity(velocity)
				px := position[0]
				py := position[1]
				if py > localMaxY {
					localMaxY = py
				}
				nx := nextVelocity[0]
				ny := nextVelocity[1]
				nextPosition := []int{px + nx, py + ny}
				position = nextPosition
				velocity = nextVelocity
				iterations += 1
				if onTarget {
					hit := stringifyCoordinates(initVelocity)
					uniqueHits[hit] = true
				}
				if localMaxY > maxY && onTarget {
					maxY = localMaxY
					maxInitVelocity = initVelocity
					fmt.Println("Probe landed in target area =>", onTarget, "with max y", maxY, "and initial velocity", maxInitVelocity)
				}
				if iterations >= 1000 {
					iterations = 0
					break
				}
			}
		}
	}
	fmt.Println("Part 1 - highest y position it reaches ", maxY, "at velocity:", maxInitVelocity)
	fmt.Println("Part 2 - number of distinct initial velocity values that cause the probe to be within " +
		"the target area after any step =", len(uniqueHits))
}

func calculateVelocity(position []int) []int {
	x := position[0]
	y := position[1]
	if x >= 1 {
		x -= 1
	} else if x < 0 {
		x += 1
	}
	y -= 1
	return []int{x, y}
}

func inTargetArea(position []int, targetArea map[string][]int) bool {
	onTarget := false
	x := position[0]
	y := position[1]
	xRange := targetArea["x"]
	yRange := targetArea["y"]
	if x >= xRange[0] && x <= xRange[1] {
		if y >= yRange[0] && y <= yRange[1] {
			onTarget = true
		}
	}
	return onTarget
}

func setTargetArea(puzzleInput []string) map[string][]int {
	targetArea := make(map[string][]int)
	line := strings.Split(puzzleInput[0], ":")
	parsedLine := strings.Split(line[1], ",")
	xLine := strings.Split(parsedLine[0], "=")
	yLine := strings.Split(parsedLine[1], "=")
	xParam := strings.Split(xLine[1], ".")
	yParam := strings.Split(yLine[1], ".")
	xRange := convertToIntSlice(xParam)
	yRange := convertToIntSlice(yParam)
	targetArea["x"] = xRange
	targetArea["y"] = yRange
	return targetArea
}

func convertToIntSlice(strSlice []string) []int {
	var intSlice []int
	for _, v := range strSlice {
		if v != "" {
			converted, _ := strconv.Atoi(v)
			intSlice = append(intSlice, converted)
		}
	}
	return intSlice
}