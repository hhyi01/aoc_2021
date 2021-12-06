package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day06Part1and2(filePath string)  {
	puzzleInput := strings.Split(readFileLines(filePath)[0], ",")
	lanternFishStates := make(map[int]int)
	// initialize map
	for i := 0; i <= 8; i++ {
		lanternFishStates[i] = 0
	}
	// populate initial lantern fish counts
	for _, fish := range puzzleInput {
		timer, _ := strconv.Atoi(fish)
		lanternFishStates[timer] += 1
	}
	day := 1
	totalDays := 80
	resetFish, newFish, currNum, nextNum := 0, 0, 0, 0
	for {
		fmt.Println("Current day: ", day)
		for i := 8; i >=0; i-- {
			currNum = lanternFishStates[i]
			lanternFishStates[i] = nextNum
			if i == 8 {
				lanternFishStates[i] = newFish
			}
			if i == 6 {
				lanternFishStates[i] += resetFish
			}
			nextNum = currNum
		}
		day += 1
		resetFish = lanternFishStates[0]
		newFish = lanternFishStates[0]
		if day > totalDays {
			break
		}
	}
	totalFish := sumFish(lanternFishStates)
	fmt.Println("Part 1 and 2 - total fish after", totalDays, "days =", totalFish)
}

func sumFish(fishStates map[int]int) int {
	total := 0
	for _, v := range fishStates {
		total += v
	}
	return total
}