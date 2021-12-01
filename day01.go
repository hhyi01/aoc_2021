package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func day01Part1(filePath string)  {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	prevNum := 0
	numInc := 0
	for scanner.Scan() {
		currNum, _ := strconv.Atoi(scanner.Text())
		if prevNum != 0 {
			if currNum > prevNum {
				numInc += 1
			}
		}
		prevNum = currNum
	}
	fmt.Println(numInc)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func day01Part2(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []int
	for scanner.Scan() {
		currNum, _ := strconv.Atoi(scanner.Text())
		input = append(input, currNum)
	}

	prevSum := 0
	currSum := 0
	numInc := 0
	for i, v := range input {
		if i < 3 {
			currSum += v
		} else {
			if currSum > prevSum {
				if prevSum != 0 {
					numInc += 1
				}
			}
			prevSum = currSum
			currSum -= input[i-3]
			currSum += v
		}
	}
	if currSum > prevSum {
		numInc += 1
	}
	fmt.Println(numInc)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
