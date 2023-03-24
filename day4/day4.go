package day4

import (
	"strconv"
	"strings"
)

func Part1(input string) int {
	numbers, boards := parseInput(input)

	for _, n := range numbers {
		var winningBoard *board = nil

		for _, board := range boards {
			board.mark(n)

			if board.isWin() {
				winningBoard = board
				break
			}
		}

		if winningBoard != nil {
			return winningBoard.sumUnmarkedNumbers() * n
		}
	}

	return 0
}

func Part2(input string) int {
	numbers, boards := parseInput(input)

	var lastWinningBoard *board = nil
	lastWinningNumber := -1

	for _, n := range numbers {
		for _, board := range boards {
			if !board.isWin() {
				board.mark(n)

				if board.isWin() {
					lastWinningBoard = board
					lastWinningNumber = n
				}
			}
		}
	}

	return lastWinningBoard.sumUnmarkedNumbers() * lastWinningNumber
}

func parseInput(input string) ([]int, []*board) {
	segments := strings.Split(input, "\n\n")

	numbers := make([]int, 0)
	for _, str := range strings.Split(segments[0], ",") {
		n, _ := strconv.Atoi(str)
		numbers = append(numbers, n)
	}

	boardsInput := segments[1:]
	boards := make([]*board, len(boardsInput))
	for i, boardLines := range boardsInput {
		boards[i] = newBoard(boardLines)
	}

	return numbers, boards
}
