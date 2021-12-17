package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func day15Part1(filePath string) {
	puzzleInput := readFileLines(filePath)
	_, lowestRiskPath := findPathsThruChitons(puzzleInput)
	desty := strconv.Itoa(len(puzzleInput)-1)
	destx := strconv.Itoa(len(puzzleInput[0])-1)
	dest := desty + "," + destx
	fmt.Println(dest)
	//path, riskTotal := calcRiskLevel(previousSquares, lowestRiskPath, "0,0", dest)
	fmt.Println("Part 1 - lowest total risk of any path from top left to bottom right =", lowestRiskPath[dest])
}

func day15Part2(filePath string) {
	puzzleInput := readFileLines(filePath)
	fullGrid := setGrid(len(puzzleInput)*5-1, len(puzzleInput[0])*5-1)
	fullMap := extendGrid(puzzleInput, fullGrid)
	desty := strconv.Itoa(len(fullMap)-1)
	destx := strconv.Itoa(len(fullMap[0])-1)
	dest := desty + "," + destx
	fmt.Println(dest)
	_, lowestRiskPath := findPathsThruChitons(fullMap)
	//path, riskTotal := calcRiskLevel(previousSquares, lowestRiskPath, "0,0", dest)
	fmt.Println("Part 2 - lowest total risk of any path from top left to bottom right =", lowestRiskPath[dest])
}

func extendGrid(puzzleInput []string, fullGrid [][]int) []string {
	var fullMap []string
	for y, line := range fullGrid {
		for x := range line {
			//fmt.Println("current y and x:", y, x)
			if y <= len(puzzleInput)-1 && x <= len(puzzleInput[0])-1 {
				risk, _ := strconv.Atoi(string(puzzleInput[y][x]))
				fullGrid[y][x] = risk
			} else {
				prevY := y
				prevX := x
				if y >= len(puzzleInput) {
					prevY = y - len(puzzleInput)
				}
				if y <= len(puzzleInput)-1 && x >= len(puzzleInput[0]) {
					prevX = x - len(puzzleInput[0])
				}
				newVal := fullGrid[prevY][prevX]
				if newVal == 9 {
					newVal = 1
				} else {
					newVal += 1
				}
				fullGrid[y][x] = newVal
			}
		}
	}
	for _, line := range fullGrid {
		var lineStr []string
		for _, v := range line {
			vStr := strconv.Itoa(v)
			lineStr = append(lineStr, vStr)
		}
		strLine := strings.Join(lineStr, "")
		fullMap = append(fullMap, strLine)
	}
	return fullMap
}

func findPathsThruChitons(puzzleInput []string) (map[string]string, map[string]int) {
	unvisited := make(map[string]bool)
	previousSquares := make(map[string]string)
	for y, line := range puzzleInput {
		for x := range line {
			coord := stringifyCoordinates([]int{y, x})
			unvisited[coord] = true
		}
	}
	path := make(map[string]int)
	for k := range unvisited {
		path[k] = math.MaxInt
	}
	start := "0,0"
	path[start] = 0
	for {
		currSquare := ""
		for sq := range unvisited {
			if currSquare == "" {
				currSquare = sq
			} else if path[sq] < path[currSquare] {
				currSquare = sq
			}
		}
		y, x := convertStringCoordinates(currSquare)
		neighbors := getNeighborSquares(y, x, puzzleInput)
		for _, neighbor := range neighbors {
			ny := neighbor[0]
			nx := neighbor[1]
			risk := getRisk(ny, nx, puzzleInput)
			candidate := path[currSquare] + risk
			nStr := stringifyCoordinates([]int{ny, nx})
			if candidate < path[nStr] {
				path[nStr] = candidate
				previousSquares[nStr] = currSquare
			}
		}
		delete(unvisited, currSquare)
		if len(unvisited) % 10000 == 0 {
			fmt.Println("Unvisted remaining:", len(unvisited))
		}
		if len(unvisited) == 0 {
			break
		}
	}
	return previousSquares, path
}

func calcRiskLevel(prevSquares map[string]string, path map[string]int, start string, dest string) (string, int) {
	var assembledPath []string
	square := dest
	for {
		assembledPath = append(assembledPath, square)
		square = prevSquares[square]
		if square == start {
			break
		}
	}
	var reversePath []string
	end := len(assembledPath)-1
	for i := end; i > 0; i-- {
		reversePath = append(reversePath, assembledPath[i])
	}
	return strings.Join(reversePath, " "), path[dest]
}

func getRisk(y int, x int, puzzleInput []string) int {
	risk, _ := strconv.Atoi(string(puzzleInput[y][x]))
	return risk
}

func convertStringCoordinates(coordinateStr string) (int, int) {
	c := strings.Split(coordinateStr, ",")
	y, _ := strconv.Atoi(c[0])
	x, _ := strconv.Atoi(c[1])
	return y, x
}

func stringifyCoordinates(coordinates []int) string {
	y := strconv.Itoa(coordinates[0])
	x := strconv.Itoa(coordinates[1])
	coordStr := strings.Join([]string{y,x}, ",")
	return coordStr
}

func getNeighborSquares(y int, x int, puzzleInput []string) [][]int {
	// neighbors up, down, left, right
	var neighbors [][]int
	yBounds := len(puzzleInput)
	xBounds := len(puzzleInput[0])
	up := y-1
	down := y+1
	left := x-1
	right := x+1
	if up >= 0 {
		neighbors = append(neighbors, []int{up, x})
	}
	if down < yBounds {
		neighbors = append(neighbors, []int{down, x})
	}
	if left >= 0 {
		neighbors = append(neighbors, []int{y, left})
	}
	if right < xBounds {
		neighbors = append(neighbors, []int{y, right})
	}
	return neighbors
}
