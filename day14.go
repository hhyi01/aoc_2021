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
	steps := 10
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
	fmt.Printf("Part 1 - quantity of most common element less least common after %d steps = %d", steps, mostCommon-leastCommon)
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
	steps := 40
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
		// set counts to pairs counted this step
		for k, v := range tempCounts {
			pairCounts[k] = v
		}
		// reset temp counts to 0
		for prop := range tempCounts {
			tempCounts[prop] = 0
		}
		step += 1
		if step == steps {
			break
		}
	}
	elementCounts := countElementsBasedOnPairs(pairCounts)
	fmt.Println(elementCounts)
	leastCommon, mostCommon := findMinMaxPart2(elementCounts)
	fmt.Printf("Part 2 - quantity of most common element less least common after %d steps = %d", steps, mostCommon-leastCommon)
}

func findMinMaxPart2(elementCounts map[string]int) (int, int) {
	minE := math.MaxInt
	maxE := math.MinInt
	for _, v := range elementCounts {
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
	for c, count := range elementCounts {
		if count % 2 == 1 {
			elementCounts[c] = (count+1)/2
		} else {
			elementCounts[c] = count/2
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