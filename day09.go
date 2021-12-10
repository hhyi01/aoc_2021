package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func day09Part1(filePath string)  {
	puzzleInput := readFileLines(filePath)
	// find low points
	var lowPoints []int
	for y, heightMap := range puzzleInput {
		for x, height := range heightMap {
			currHeight, _ := strconv.Atoi(string(height))
			neighbors := findNeighbors(y, x, puzzleInput, heightMap)
			// compare with neighbors
			isLowPoint := false
			minNeighbor := math.MaxInt
			// find min neighbor
			for _, neighbor := range neighbors {
				if neighbor < minNeighbor {
					minNeighbor = neighbor
				}
			}
			if currHeight < minNeighbor {
				isLowPoint = true
			}
			if isLowPoint {
				lowPoints = append(lowPoints, currHeight)
				fmt.Println("Found a low point:", currHeight, neighbors)
			}
		}
	}
	riskLevelTotal := 0
	for _, l := range lowPoints {
		riskLevelTotal += l + 1
	}
	fmt.Println("Part 1 - sum of the risk levels of all low points on heighmap =", riskLevelTotal)
}

func day09Part2(filePath string)  {
	puzzleInput := readFileLines(filePath)
	var basins []int
	dimy := len(puzzleInput)
	dimx := len(puzzleInput[0])
	visited := setGridBool(dimx, dimy)
	neighborsChecked := setGridBool(dimx, dimy)
	for y, heightMap := range puzzleInput {
		for x, height := range heightMap {
			currHeight, _ := strconv.Atoi(string(height))
			neighbors := findNeighbors(y, x, puzzleInput, heightMap)
			// compare with neighbors
			isLowPoint := false
			minNeighbor := math.MaxInt
			// find min neighbor
			for _, neighbor := range neighbors {
				if neighbor < minNeighbor {
					minNeighbor = neighbor
				}
			}
			if currHeight < minNeighbor {
				isLowPoint = true
			}
			if isLowPoint {
				var basin []int
				// find neighbors starting with low point
				basin = append(basin, currHeight)
				// visited low point
				visited[y][x] = true
				// queue neighbors to check
				neighborsToCheck := findNeighborsPart2(y, x, puzzleInput, heightMap)
				neighborsChecked[y][x] = true
				for {
					currNeighbor := neighborsToCheck[len(neighborsToCheck)-1]
					currY := currNeighbor[0]
					currX := currNeighbor[1]
					currNeighborHeight, _ := strconv.Atoi(string(puzzleInput[currY][currX]))
					isVisited := visited[currY][currX]
					if !isVisited {
						if currNeighborHeight != 9 {
							basin = append(basin, currNeighborHeight)
						}
						visited[currY][currX] = true
					}
					// remove current neighbor from queue
					neighborsToCheck = neighborsToCheck[:len(neighborsToCheck)-1]
					foundNeighbors := neighborsChecked[currY][currX]
					if !foundNeighbors {
						if currNeighborHeight != 9 {
							nextNeighbors := findNeighborsPart2(currY, currX, puzzleInput, heightMap)
							neighborsToCheck = append(neighborsToCheck, nextNeighbors...)
						}
						neighborsChecked[currY][currX] = true
					}
					if len(neighborsToCheck) == 0 {
						break
					}
				}
				basins = append(basins, len(basin))
			}
		}
	}
	sort.Ints(basins)
	fmt.Println("Part 2 - product of sizes of 3 largest basins =",
		basins[len(basins)-1]*basins[len(basins)-2]*basins[len(basins)-3])
}

func setGridBool(x int, y int) [][]bool {
	grid := make([][]bool, y+1)
	for i := range grid {
		grid[i] = make([]bool, x+1)
	}
	return grid
}

func findNeighbors(y int, x int, puzzleInput []string, heightMap string) []int {
	var neighbors []int
	up := y-1
	down := y+1
	left := x-1
	right := x+1
	if up >= 0 {
		upHeight, _ := strconv.Atoi(string(puzzleInput[up][x]))
		neighbors = append(neighbors, upHeight)
	}
	if down < len(puzzleInput) {
		downHeight, _ := strconv.Atoi(string(puzzleInput[down][x]))
		neighbors = append(neighbors, downHeight)
	}
	if left >= 0 {
		leftHeight, _ := strconv.Atoi(string(puzzleInput[y][left]))
		neighbors = append(neighbors, leftHeight)
	}
	if right < len(heightMap) {
		rightHeight, _ := strconv.Atoi(string(puzzleInput[y][right]))
		neighbors = append(neighbors, rightHeight)
	}
	return neighbors
}

func findNeighborsPart2(y int, x int, puzzleInput []string, heightMap string) [][]int {
	var neighbors [][]int
	up := y-1
	down := y+1
	left := x-1
	right := x+1
	if up >= 0 {
		upCoord := []int{up, x}
		neighbors = append(neighbors, upCoord)
	}
	if down < len(puzzleInput) {
		downCoord := []int{down, x}
		neighbors = append(neighbors, downCoord)
	}
	if left >= 0 {
		leftCoord := []int{y, left}
		neighbors = append(neighbors, leftCoord)
	}
	if right < len(heightMap) {
		rightCoord := []int{y, right}
		neighbors = append(neighbors, rightCoord)
	}
	return neighbors
}