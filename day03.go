package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func day03Part1(filePath string)  {
	puzzleInput := readFileLines(filePath)
	// only counting 0s
	counts := initializeCounts(puzzleInput)
	// set epsilon and gamma
	inputLength := len(puzzleInput)
	epsilon, gamma := setRatesBinary(inputLength, counts)
	epsilonRate := convertToDecimal(epsilon)
	gammaRate := convertToDecimal(gamma)
	fmt.Println("Part 1 - power consumption of submarine = ", epsilonRate * gammaRate)
}

func convertToDecimal(binaryStr string) int64 {
	rate, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return rate
}

func setRatesBinary(inputLength int, counts []int) (string, string) {
	var eTemp bytes.Buffer
	var gTemp bytes.Buffer
	for _, v := range counts {
		if v > inputLength/2 {
			eTemp.WriteString("0")
			gTemp.WriteString("1")
		} else {
			eTemp.WriteString("1")
			gTemp.WriteString("0")
		}
	}
	epsilon := eTemp.String()
	gamma := gTemp.String()
	return epsilon, gamma
}

func initializeCounts(puzzleInput []string) []int {
	var counts []int
	numPos := len(puzzleInput[0])
	for i := 0; i < numPos; i++ {
		counts = append(counts, 0)
	}
	for _, v := range puzzleInput {
		for pos, b := range v {
			currBit := string(b)
			if currBit == "0" {
				counts[pos] += 1
			}
		}
	}
	return counts
}

func day03Part2(filePath string) {
	puzzleInput := readFileLines(filePath)
	// set O2 and CO2
	o2GeneratorRating := calcRating(puzzleInput, "O2")
	co2ScrubberRating := calcRating(puzzleInput, "CO2")
	fmt.Println("Part 2 - life support rating = ", o2GeneratorRating * co2ScrubberRating)
}

func calcRating(puzzleInput []string, ratingType string) int64 {
	position := 0
	counts := initializeCounts(puzzleInput)
	var rating string
	for {
		count := counts[position]
		var binaryNums []string
		var target string
		if ratingType == "O2" {
			if count > len(puzzleInput)/2 {
				target = "0"
			} else {
				target = "1"
			}
		} else {
			if count > len(puzzleInput)/2 {
				target = "1"
			} else {
				target = "0"
			}
		}
		for _, v := range puzzleInput {
			if string(v[position]) == target {
				binaryNums = append(binaryNums, v)
			}
		}
		position += 1
		counts = initializeCounts(binaryNums)
		puzzleInput = binaryNums
		if len(binaryNums) == 1 {
			rating = binaryNums[0]
			break
		}
	}
	return convertToDecimal(rating)
}
