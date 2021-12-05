package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day04Part1(filePath string) {
	puzzleInput := readFileLines(filePath)
	bingoNums, bingoCards := parseBingoInput(puzzleInput)
	var score int
	for _, v := range bingoNums {
		bingoCards = markNumbers(v, bingoCards)
		wiener := checkForCardWinner(bingoCards)
		if len(wiener) > 0 {
			for _, line := range wiener {
				fmt.Println(line)
			}
			unmarkedSum := sumUnmarkedNums(wiener)
			lastNum, _ := strconv.Atoi(v)
			score = lastNum * unmarkedSum
			break
		}
	}
	fmt.Println("Part 1 - final score = ", score)
}

func day04Part2(filePath string) {
	puzzleInput := readFileLines(filePath)
	bingoNums, bingoCards := parseBingoInput(puzzleInput)
	var winningBoards [][]int
	var currentWinners []int
	winningBoardsStates := make(map[int][][]string)
	for _, v := range bingoNums {
		bingoCards = markNumbers(v, bingoCards)
		currentWinners = checkForCardWinnerPart2(bingoCards)
		for _, board := range currentWinners {
			_, found := winningBoardsStates[board]
			if !found {
				winningBoardsStates[board] = copyCard(bingoCards[board])
			}
		}
		winningBoards = addNewWinner(winningBoards, currentWinners, v)
	}
	fmt.Println("Winning card(s) in order of wins with their winning numbers [card #, winning #]: ", winningBoards)
	lastWinner := winningBoards[len(winningBoards)-1]
	fmt.Println("Last winning board: ")
	for _, v := range winningBoardsStates[lastWinner[0]] {
		fmt.Println(v)
	}
	unmarkedSum := sumUnmarkedNums(winningBoardsStates[lastWinner[0]])
	lastNum := lastWinner[1]
	score := unmarkedSum * lastNum
	fmt.Println("Part 2 - final score = ", score)
}

func copyCard(card [][]string) [][]string {
	var savedCard [][]string
	for _, v := range card {
		var line []string
		for _, e := range v {
			line = append(line, e)
		}
		savedCard = append(savedCard, line)
	}
	return savedCard
}

func parseBingoInput(puzzleInput []string) ([]string, [][][]string) {
	var bingoNums []string
	var bingoCards [][][]string
	var card [][]string
	for i, v := range puzzleInput {
		if i == 0 {
			bingoNums = strings.Split(v, ",")
		} else {
			if v != "" {
				line := strings.Split(v, " ")
				emptyStrRemoved := removeEmptyStrings(line)
				card = append(card, emptyStrRemoved)
			} else {
				if len(card) > 0 {
					bingoCards = append(bingoCards, card)
					card = make([][]string, 0)
				}
			}
		}
	}
	bingoCards = append(bingoCards, card)
	return bingoNums, bingoCards
}

func removeEmptyStrings(line []string) []string {
	var emptyStrRemoved []string
	for _, v := range line {
		if v != "" {
			emptyStrRemoved = append(emptyStrRemoved, v)
		}
	}
	return emptyStrRemoved
}

func markNumbers(bingoNum string, bingoCards [][][]string) [][][]string {
	for i, v := range bingoCards {
		for n, c := range v {
			for m, l := range c {
				if l == bingoNum {
					bingoCards[i][n][m] = "D"
				}
			}
		}
	}
	return bingoCards
}

func checkForCardWinner(bingoCards [][][]string) [][]string {
	var winningCard [][]string
	for _, card := range bingoCards {
		hasWinningRow := checkRowWinner(card)
		hasWinningColumn := checkColumnWinner(card)
		if hasWinningRow || hasWinningColumn {
			winningCard = card
			break
		}
	}
	return winningCard
}

func checkForCardWinnerPart2(bingoCards [][][]string) []int {
	var winningCards []int
	for i, card := range bingoCards {
		hasWinningRow := checkRowWinner(card)
		hasWinningColumn := checkColumnWinner(card)
		if hasWinningRow || hasWinningColumn {
			winningCards = append(winningCards, i)
		}
	}
	return winningCards
}

func addNewWinner(winningBoards [][]int, currentWinners []int, currentNum string) [][]int {
	for _, winner := range currentWinners {
		if !contains(winningBoards, winner) {
			winningNum, _ := strconv.Atoi(currentNum)
			winningBoard := []int{winner, winningNum}
			winningBoards = append(winningBoards, winningBoard)
		}
	}
	return winningBoards
}

func contains(winningBoards [][]int, winner int) bool {
	hasWinner := false
	// store winners like this [board number, winning number]
	for _, v := range winningBoards {
		if v[0] == winner {
			hasWinner = true
		}
	}
	return hasWinner
}

func checkRowWinner(card [][]string) bool {
	rowWinner := false
	for _, v := range card {
		marked := countRowMarked(v)
		if marked == len(v) {
			rowWinner = true
		}
	}
	return rowWinner
}

func countRowMarked(cardLine []string) int {
	marked := 0
	for _, v := range cardLine {
		if v == "D" {
			marked += 1
		}
	}
	return marked
}

func checkColumnWinner(card [][]string) bool {
	marked := 0
	idx := 0
	columnWinner := false
	for {
		for _, v := range card {
			if v[idx] == "D" {
				marked += 1
			}
		}
		if marked == len(card) {
			columnWinner = true
			break
		}
		idx += 1
		marked = 0
		if idx >= len(card) {
			break
		}
	}
	return columnWinner
}

func sumUnmarkedNums(card [][]string) int {
	unmarkedSum := 0
	for _, line := range card {
		for _, v := range line {
			n, _ := strconv.Atoi(v)
			unmarkedSum += n
		}
	}
	return unmarkedSum
}
