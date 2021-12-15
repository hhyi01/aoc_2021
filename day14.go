package main

import (
	"fmt"
	"math"
	"strings"
)

func day14Part1(filePath string) {
	puzzleInput := readFileLines(filePath)
	polymerTemplate, rules := parseInput(puzzleInput)
	step := 0
	steps := 4
	polymer := polymerTemplate
	for {
		var nextPolymer []string
		for i := 0; i < len(polymer); i++ {
			if i < len(polymer)-1 {
				pair := polymer[i] + polymer[i+1]
				insert := rules[pair]
				nextPolymer = append(nextPolymer, polymer[i] + insert)
			} else {
				nextPolymer = append(nextPolymer, polymer[i])
			}
		}
		nextPolymerStr := strings.Join(nextPolymer, "")
		polymer = strings.Split(nextPolymerStr, "")
		step += 1
		if step == steps {
			break
		}
	}
	fmt.Println(polymer)
	leastCommon, mostCommon := findMinMax(polymer)
	fmt.Println("Part 1 - quantity of most common element less least common =", mostCommon - leastCommon)
}

func day14Part2(filePath string)  {
	puzzleInput := readFileLines(filePath)
	polymerTemplate, rules := parseInput(puzzleInput)
	pairCounts := make(map[string]int)
	tempCounts := make(map[string]int)
	for pair := range rules {
		pairCounts[pair] = 0
	}
	// init pairs map
	for i := 0; i < len(polymerTemplate)-1; i++ {
		pair := polymerTemplate[i] + polymerTemplate[i+1]
		pairCounts[pair] += 1
		tempCounts[pair] = 0
	}
	step := 0
	steps := 1
	for {
		for pair, count := range pairCounts {
			if count > 0 {
				insert := rules[pair]
				newPair1 := string(pair[0]) + insert
				newPair2 := insert + string(pair[1])
				tempCounts[newPair1] += count
				tempCounts[newPair2] += count
			}
		}
		for k, v := range tempCounts {
			pairCounts[k] = v
		}
		// reset counts to 0
		for prop := range tempCounts {
			tempCounts[prop] = 0
		}
		step += 1
		if step == steps {
			break
		}
	}
	fmt.Println(pairCounts)
	elementCounts := countElementsBasedOnPairs(pairCounts)
	fmt.Println(elementCounts)
	ans := "NCNBCHB"
	charCounts := countString(ans)
	fmt.Println(charCounts)
	fmt.Println("Part 2")
}

func countString(str string) map[string]int {
	charCounts := make(map[string]int)
	for _, v := range str {
		_, found := charCounts[string(v)]
		if found {
			charCounts[string(v)] += 1
		} else {
			charCounts[string(v)] = 1
		}
	}
	return charCounts
}

func countElementsBasedOnPairs(pairCounts map[string]int) map[string]int {
	elementCounts := make(map[string]int)
	for k, v := range pairCounts {
		pair := strings.Split(k, "")
		e1 := pair[0]
		e2 := pair[1]
		_, found1 := elementCounts[e1]
		_, found2 := elementCounts[e2]
		if found1 {
			elementCounts[e1] += v
		} else {
			elementCounts[e1] = v
		}
		if found2 {
			elementCounts[e2] += v
		} else {
			elementCounts[e2] = v
		}
	}
	return elementCounts
}

func findMinMax(polymer []string) (int, int) {
	minE := math.MaxInt
	maxE := math.MinInt
	countE := make(map[string]int)
	for _, e := range polymer {
		_, found := countE[e]
		if found {
			countE[e] += 1
		} else {
			countE[e] = 1
		}
	}
	fmt.Println(countE)
	for _, v := range countE {
		if v > maxE {
			maxE = v
		} else {
			if v < minE {
				minE = v
			}
		}
	}
	return minE, maxE
}

func parseInput(puzzleInput []string) ([]string, map[string]string) {
	var polymerTemplate []string
	rules := make(map[string]string)
	for i, line := range puzzleInput {
		if i == 0 {
			polymerTemplate = strings.Split(puzzleInput[i],"")
		} else {
			if line != "" {
				rule := strings.Split(line, " ")
				pair := rule[0]
				insert := rule[2]
				rules[pair] = insert
			}
		}
	}
	return polymerTemplate, rules
}