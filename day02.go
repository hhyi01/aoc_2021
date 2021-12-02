package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func day02Part1(filePath string)  {
	puzzleInput := readFileLines(filePath)
	horizontal, depth := 0, 0
	for _, line := range puzzleInput {
		command := strings.Split(line, " ")
		direction := command[0]
		units, _ := strconv.Atoi(command[1])
		horizontal, depth = day02Part1CalcPosition(direction, units, horizontal, depth)
	}
	fmt.Println("Part 1 - final horizontal x final depth = ", horizontal * depth)
}

func day02Part1CalcPosition(direction string, units int, horizontal int, depth int) (int, int) {
	switch direction {
	case "forward":
		horizontal += units
	case "down":
		depth += units
	case "up":
		depth -= units
	default:
		log.Fatalf("unknown direction %q", direction)
	}
	return horizontal, depth
}

func day02Part2(filePath string)  {
	puzzleInput := readFileLines(filePath)
	horizontal, depth, aim := 0, 0, 0
	for _, line := range puzzleInput {
		command := strings.Split(line, " ")
		direction := command[0]
		units, _ := strconv.Atoi(command[1])
		horizontal, depth, aim = day02Part2CalcPosition(direction, units, horizontal, depth, aim)
	}
	fmt.Println("Part 2 - final horizontal x final depth = ", horizontal * depth)
}

func day02Part2CalcPosition(direction string, units int, horizontal int, depth int, aim int) (int, int, int) {
	switch direction {
	case "forward":
		horizontal += units
		depth += aim * units
	case "down":
		aim += units
	case "up":
		aim -= units
	default:
		log.Fatalf("unknown direction %q", direction)
	}
	return horizontal, depth, aim
}

