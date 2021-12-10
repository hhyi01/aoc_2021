package main

import (
	"fmt"
	"sort"
	"strings"
)

func day10Part1(filePath string)  {
	puzzleInput := readFileLines(filePath)
	illegalCharsCount := map[string]int{")":0, "]":0, "}":0, ">":0}
	matchingChars := map[string]string{")":"(", "]":"[", "}":"{", ">":"<"}
	scoreChars := map[string]int{")":3, "]":57, "}":1197, ">":25137}
	for _, v := range puzzleInput {
		line := strings.Split(v, "")
		var chunks []string
		for _, c := range line {
			if c == "(" || c == "[" || c == "{" || c == "<" {
				chunks = append(chunks, c)
			} else {
				if len(chunks) > 0 {
					lastChar := chunks[len(chunks)-1]
					chunks = chunks[:len(chunks)-1]
					legalChar, _ := matchingChars[c]
					if lastChar != legalChar {
						illegalCharsCount[c] += 1
					}
				} else {
					illegalCharsCount[c] += 1
				}
			}
		}
		fmt.Println("Current chunk:", chunks)
	}
	fmt.Println(illegalCharsCount)
	errorScore := 0
	for k, v := range illegalCharsCount {
		points := scoreChars[k]
		errorScore += v * points
	}
	fmt.Println("Part 1 - total syntax error score =", errorScore)
}

func day10Part2(filePath string)  {
	puzzleInput := readFileLines(filePath)
	matchingChars := map[string]string{")":"(", "]":"[", "}":"{", ">":"<", "(":")", "[":"]", "{":"}", "<":">"}
	scoreChars := map[string]int{")":1, "]":2, "}":3, ">":4}
	var scores []int
	for _, v := range puzzleInput {
		line := strings.Split(v, "")
		var chunks []string
		illegalCharsCount := 0
		for _, c := range line {
			if c == "(" || c == "[" || c == "{" || c == "<" {
				chunks = append(chunks, c)
			} else {
				if len(chunks) > 0 {
					lastChar := chunks[len(chunks)-1]
					chunks = chunks[:len(chunks)-1]
					legalChar, _ := matchingChars[c]
					if lastChar != legalChar {
						illegalCharsCount += 1
					}
				} else {
					illegalCharsCount += 1
				}
			}
		}
		if illegalCharsCount > 0 {
			fmt.Println("Corrupted line - disregard.")
		} else {
			fmt.Println("Current chunk:", chunks)
			closingSeq := assembleClosingSeq(chunks, matchingChars)
			score := calcScore(closingSeq, scoreChars)
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	fmt.Println("Part 2 - total syntax error score =", scores[len(scores)/2])
}

func assembleClosingSeq(chunks []string, matchingChars map[string]string) []string {
	var closingSeq []string
	idx := len(chunks)-1
	for {
		currChar := chunks[idx]
		missingChar, _ := matchingChars[currChar]
		closingSeq = append(closingSeq, missingChar)
		idx -= 1
		if idx < 0 {
			break
		}
	}
	return closingSeq
}

func calcScore(closingSeq []string, scoreChars map[string]int) int {
	score := 0
	for _, c := range closingSeq {
		score = 5 * score
		score += scoreChars[c]
	}
	return score
}