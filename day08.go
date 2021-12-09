package main

import (
	"fmt"
	"strings"
)

func day08Part1(filePath string)  {
	segmentCount := make(map[int]int)
	segmentCount[2] = 1
	segmentCount[4] = 4
	segmentCount[3] = 7
	segmentCount[7] = 8
	countNums := 0
	puzzleInput := readFileLines(filePath)
	for _, p := range puzzleInput {
		line := strings.Split(p, "|")
		digitalOutput := line[1]
		digits := strings.Split(digitalOutput, " ")
		for _, d := range digits {
			digitStr := strings.TrimSpace(d)
			digitLen := len(digitStr)
			_, found := segmentCount[digitLen]
			if found {
				countNums += 1
			}
		}
	}
	fmt.Println("Part 1 - number of times 1, 4, 7, or 8 appear =", countNums)
}
