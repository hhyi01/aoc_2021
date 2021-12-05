package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day05Part1(filePath string)  {
	puzzleInput := readFileLines(filePath)
	maxX, maxY := getMaxXandY(puzzleInput)
	fmt.Println("Max x and y: ", maxX, maxY)
	grid := setGrid(maxX, maxY)
	grid = markVentsPart1(puzzleInput, grid)
	overlapCount := countOverlapPoints(grid)
	fmt.Println("Part 1 - number of points where at least 2 lines overlap = ", overlapCount)
}

func day05Part2(filePath string)  {
	puzzleInput := readFileLines(filePath)
	maxX, maxY := getMaxXandY(puzzleInput)
	fmt.Println("Max x and y: ", maxX, maxY)
	grid := setGrid(maxX, maxY)
	grid = markVentsPart2(puzzleInput, grid)
	overlapCount := countOverlapPoints(grid)
	fmt.Println("Part 2 - number of points where at least 2 lines overlap = ", overlapCount)
}

func countOverlapPoints(grid [][]int) int {
	overlapCount := 0
	for _, line := range grid {
		for _, v := range line {
			if v > 1 {
				overlapCount += 1
			}
		}
	}
	return overlapCount
}

func getXandYValues(line string) (int, int, int, int) {
	parsedLine := strings.Split(line, " ")
	startCoord := strings.Split(parsedLine[0], ",")
	endCoord := strings.Split(parsedLine[len(parsedLine)-1], ",")
	x1, _ := strconv.Atoi(startCoord[0])
	x2, _ := strconv.Atoi(endCoord[0])
	y1, _ := strconv.Atoi(startCoord[1])
	y2, _ := strconv.Atoi(endCoord[1])
	return x1, x2, y1, y2
}

func markVentsPart1(puzzleInput []string, grid [][]int) [][]int {
	for _, v := range puzzleInput {
		x1, x2, y1, y2 := getXandYValues(v)
		minX := calcMin(x1, x2)
		maxX := calcMax(x1, x2)
		minY := calcMin(y1, y2)
		maxY := calcMax(y1, y2)
		if x1 != x2 && y1 == y2 {
			for i := minX; i <= maxX; i++ {
				grid[y1][i] += 1
			}
		}
		if y1 != y2 && x1 == x2 {
			for i := minY; i <= maxY; i++ {
				grid[i][x1] += 1
			}
		}
	}
	return grid
}

func markVentsPart2(puzzleInput []string, grid [][]int) [][]int {
	for _, v := range puzzleInput {
		x1, x2, y1, y2 := getXandYValues(v)
		minX := calcMin(x1, x2)
		maxX := calcMax(x1, x2)
		minY := calcMin(y1, y2)
		maxY := calcMax(y1, y2)
		if x1 != x2 && y1 == y2 {
			for i := minX; i <= maxX; i++ {
				grid[y1][i] += 1
			}
		}
		if y1 != y2 && x1 == x2 {
			for i := minY; i <= maxY; i++ {
				grid[i][x1] += 1
			}
		}
		if x1 != x2 && y1 != y2 {
			incX := true
			incY := true
			if x1 > x2 {
				incX = false
			}
			if y1 > y2 {
				incY = false
			}
			xCoord := x1
			yCoord := y1
			for {
				grid[yCoord][xCoord] += 1
				if incX {
					xCoord += 1
				} else {
					xCoord -= 1
				}
				if incY {
					yCoord += 1
				} else {
					yCoord -= 1
				}
				if xCoord == x2 || yCoord == y2 {
					break
				}
			}
			grid[yCoord][xCoord] += 1
		}
	}
	return grid
}

func getMaxXandY(puzzleInput []string) (int, int) {
	maxX := 0
	maxY := 0
	for _, v := range puzzleInput {
		x1, x2, y1, y2 := getXandYValues(v)
		maxXCoord := calcMax(x1, x2)
		maxYCoord := calcMax(y1, y2)
		if maxXCoord > maxX {
			maxX = maxXCoord
		}
		if maxYCoord > maxY {
			maxY = maxYCoord
		}
	}
	return maxX, maxY
}

func calcMax(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func calcMin(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}

func setGrid(x int, y int) [][]int {
	grid := make([][]int, y+1)
	for i := range grid {
		grid[i] = make([]int, x+1)
	}
	return grid
}