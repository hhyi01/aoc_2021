package main

import (
	"fmt"
	"strconv"
)

func day11Part1(filePath string)  {
	puzzleInput := readFileLines(filePath)
	octopusGrid := setGrid(len(puzzleInput[0])-1, len(puzzleInput)-1)
	for i, line := range puzzleInput {
		for n, o := range line {
			octopus, _ := strconv.Atoi(string(o))
			octopusGrid[i][n] = octopus
		}
	}
	// keep track of who flashed this step
	flashed := setGridBool(len(puzzleInput[0])-1, len(puzzleInput)-1)
	steps := 100
	step := 1
	flashedCount := 0
	for {
		fmt.Println("State of octopodes at step", step)
		printGrid(octopusGrid)
		octopusGrid = addEnergyWholeGrid(octopusGrid)
		flashers := findFlashers(octopusGrid)
		fmt.Println("Current flashers:", flashers)
		var neighbors [][]int
		for {
			for _, flasher := range flashers {
				foundNeighbors := findFlasherNeighbors(flasher, octopusGrid)
				neighbors = append(neighbors, foundNeighbors...)
				y := flasher[0]
				x := flasher[1]
				octopusGrid[y][x] = 0
				flashed[y][x] = true
			}
			octopusGrid = addEnergyToNeighbors(neighbors, octopusGrid, flashed)
			neighbors = neighbors[:0]
			flashers = findFlashers(octopusGrid)
			if len(flashers) == 0 {
				break
			}
		}
		flashedCount += countFlashes(flashed)
		step += 1
		flashed = setGridBool(len(puzzleInput[0])-1, len(puzzleInput)-1)
		if step > steps {
			break
		}
	}
	fmt.Println("Part 1 - total flashes after", steps, "steps =", flashedCount)
}

func day11Part2(filePath string) {
	puzzleInput := readFileLines(filePath)
	octopusGrid := setGrid(len(puzzleInput[0])-1, len(puzzleInput)-1)
	for i, line := range puzzleInput {
		for n, o := range line {
			octopus, _ := strconv.Atoi(string(o))
			octopusGrid[i][n] = octopus
		}
	}
	// keep track of who flashed this step
	flashed := setGridBool(len(puzzleInput[0])-1, len(puzzleInput)-1)
	step := 0
	flashedCount := 0
	totalOctopodes := len(puzzleInput[0]) * len(puzzleInput)
	fmt.Println("Total octopodes:", totalOctopodes)
	for {
		fmt.Println("State of octopodes at step", step)
		printGrid(octopusGrid)
		octopusGrid = addEnergyWholeGrid(octopusGrid)
		flashers := findFlashers(octopusGrid)
		fmt.Println("Current flashers:", flashers)
		var neighbors [][]int
		for {
			for _, flasher := range flashers {
				foundNeighbors := findFlasherNeighbors(flasher, octopusGrid)
				neighbors = append(neighbors, foundNeighbors...)
				y := flasher[0]
				x := flasher[1]
				octopusGrid[y][x] = 0
				flashed[y][x] = true
			}
			octopusGrid = addEnergyToNeighbors(neighbors, octopusGrid, flashed)
			neighbors = neighbors[:0]
			flashers = findFlashers(octopusGrid)
			if len(flashers) == 0 {
				break
			}
		}
		flashedCount = countFlashes(flashed)
		step += 1
		flashed = setGridBool(len(puzzleInput[0])-1, len(puzzleInput)-1)
		if flashedCount == totalOctopodes {
			break
		}
	}
	fmt.Println("Part 1 - first step where all octopodes flash =", step)
}

func findFlasherNeighbors(flasher []int, octopusGrid [][]int) [][]int {
	y := flasher[0]
	x := flasher[1]
	up := y-1
	down := y+1
	left := x-1
	right := x+1
	var neighbors [][]int
	if up >= 0 {
		// find diagonal left and right also
		upNeighbor := []int{up, x}
		neighbors = append(neighbors, upNeighbor)
		var diagUpLeft []int
		var diagUpRight []int
		if left >= 0 {
			diagUpLeft = append(diagUpLeft, up, left)
			neighbors = append(neighbors, diagUpLeft)
		}
		if right < len(octopusGrid[0]) {
			diagUpRight = append(diagUpRight, up, right)
			neighbors = append(neighbors, diagUpRight)
		}
	}
	if down < len(octopusGrid) {
		downNeighbor := []int{down, x}
		neighbors = append(neighbors, downNeighbor)
		var diagDownLeft []int
		var diagDownRight []int
		if left >= 0 {
			diagDownLeft = append(diagDownLeft, down, left)
			neighbors = append(neighbors, diagDownLeft)
		}
		if right < len(octopusGrid[0]) {
			diagDownRight = append(diagDownRight, down, right)
			neighbors = append(neighbors, diagDownRight)
		}
	}
	if left >= 0 {
		neighbors = append(neighbors, []int{y, left})
	}
	if right < len(octopusGrid[0]) {
		neighbors = append(neighbors, []int{y, right})
	}
	return neighbors
}

func addEnergyWholeGrid(octopusGrid [][]int) [][]int {
	for y, line := range octopusGrid {
		for x := range line {
			octopusGrid[y][x] += 1
		}
	}
	return octopusGrid
}

func addEnergyToNeighbors(neighbors [][]int, octopusGrid [][]int, flashed [][]bool) [][]int {
	for _, neighbor := range neighbors {
		y := neighbor[0]
		x := neighbor[1]
		flashedThisStep := flashed[y][x]
		if !flashedThisStep {
			octopusGrid[y][x] += 1
		}
	}
	return octopusGrid
}

func findFlashers(octopusGrid [][]int) [][]int {
	var flashers [][]int
	for y, line := range octopusGrid {
		for x, o := range line {
			if o > 9 {
				octopus := []int{y, x}
				flashers = append(flashers, octopus)
			}
		}
	}
	return flashers
}

func printGrid(grid [][]int) {
	for _, v := range grid {
		fmt.Println(v)
	}
}

func countFlashes(flashed [][]bool) int {
	count := 0
	for _, line := range flashed {
		for _, v := range line {
			if v {
				count += 1
			}
		}
	}
	return count
}