package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day13Part1(filePath string)  {
	puzzleInput := readFileLines(filePath)
	coordinates, folds := parseCoordinatesFolds(puzzleInput)
	maxX, maxY := getMaxXandY(coordinates)
	paper := setGridString(maxX+1, maxY+1)
	paper = markDots(coordinates, paper)
	for _, line := range folds {
		// run only for first fold in Part 1
		instruction := strings.Split(line, " ")
		fold := strings.Split(instruction[2], "=")
		axis := fold[0]
		lineNum, _ := strconv.Atoi(fold[1])
		if axis == "y" {
			paper = foldVertically(paper, lineNum, coordinates)
			paper = resizePaperVertical(lineNum, paper)
		} else {
			paper = foldHorizontally(paper, lineNum, coordinates)
			paper = resizePaperHorizontal(lineNum, paper)
		}
	}
	fmt.Println("Part 2: code to activate infrared thermal imaging camera system:")
	for _, p := range paper {
		fmt.Println(p)
	}
	//dotCount := countDots(paper)
	//fmt.Println("Part 1 - dots visible after completing first fold =", dotCount)
}

func foldVertically(paper [][]string, yFold int, coordinates []string) [][]string {
	// find coordinates below fold line
	// all x-coordinates stay the same
	// shift y-coordinates
	for y, line := range paper {
		for x, e := range line {
			if e == "#" {
				if y > yFold {
					newy := (y - yFold * 2) * -1
					paper[newy][x] = "#"
				}
			}
		}
	}
	return paper
}

func foldHorizontally(paper [][]string, xFold int, coordinates []string) [][]string {
	for y, line := range paper {
		for x, e := range line {
			if e == "#" {
				if x > xFold {
					newx := (x - xFold * 2) * -1
					paper[y][newx] = "#"
				}
			}
		}
	}
	return paper
}

func resizePaperVertical(yFold int, paper [][]string) [][]string {
	var resizedPaper [][]string
	for i, line := range paper {
		if i < yFold {
			resizedPaper = append(resizedPaper, line)
		}
	}
	return resizedPaper
}

func resizePaperHorizontal(xFold int, paper [][]string) [][]string {
	var resizedPaper [][]string
	for _, line := range paper {
		resizedPaper = append(resizedPaper, line[:xFold])
	}
	return resizedPaper
}

func countDots(paper [][]string) int {
	dots := 0
	for _, line := range paper {
		for _, e := range line {
			if e == "#" {
				dots += 1
			}
		}
	}
	return dots
}

func markDots(coordinates []string, paper [][]string) [][]string {
	for i, line := range paper {
		for n := range line {
			paper[i][n] = "."
		}
	}
	for _, c := range coordinates {
		coord := strings.Split(c, ",")
		x, _ := strconv.Atoi(coord[0])
		y, _ := strconv.Atoi(coord[1])
		paper[y][x] = "#"
	}
	return paper
}

func setGridString(x int, y int) [][]string {
	grid := make([][]string, y)
	for i := range grid {
		grid[i] = make([]string, x)
	}
	return grid
}

func parseCoordinatesFolds(puzzleInput []string) ([]string, []string) {
	var coordinates []string
	var folds []string
	coord := true
	for _, v := range puzzleInput {
		if v == "" {
			coord = false
		}
		if coord {
			coordinates = append(coordinates, v)
		} else {
			if v != "" {
				folds = append(folds, strings.TrimSpace(v))
			}
		}
	}
	return coordinates, folds
}