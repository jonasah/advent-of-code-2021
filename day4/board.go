package day4

import (
	"strconv"
	"strings"
)

const boardSize = 5

type boardNumber struct {
	number int
	marked bool
}

type board struct {
	grid [boardSize][boardSize]*boardNumber
}

func newBoard(boardLines string) *board {
	b := &board{}

	for r, row := range strings.Split(boardLines, "\n") {
		columns := strings.Fields(row)

		for c, column := range columns {
			n, _ := strconv.Atoi(column)
			b.grid[r][c] = &boardNumber{number: n, marked: false}
		}
	}

	return b
}

func (b *board) mark(n int) {
	for _, row := range b.grid {
		for _, col := range row {
			if col.number == n {
				col.marked = true
				return
			}
		}
	}
}

func (b *board) isWin() bool {
	for i := 0; i < boardSize; i++ {
		rowCompleted := true
		columnCompleted := true

		for j := 0; j < boardSize; j++ {
			if !b.grid[i][j].marked {
				rowCompleted = false
			}
			if !b.grid[j][i].marked {
				columnCompleted = false
			}
		}

		if rowCompleted || columnCompleted {
			return true
		}
	}

	return false
}

func (b *board) sumUnmarkedNumbers() int {
	sum := 0
	for _, row := range b.grid {
		for _, col := range row {
			if !col.marked {
				sum += col.number
			}
		}
	}
	return sum
}
