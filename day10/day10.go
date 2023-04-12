package day10

import (
	"sort"

	"github.com/jonasah/advent-of-code-2021/lib/common"
	"github.com/jonasah/advent-of-code-2021/lib/container"
	"github.com/jonasah/advent-of-code-2021/lib/math"
	"github.com/jonasah/advent-of-code-2021/lib/slice"
)

var openingChars = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var closingChars = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var illegalCharPointMap = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var missingCharPointMap = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func Part1(input string) int {
	lines := common.GetLines(input)

	illegalChars := make([]rune, 0)
	for _, line := range lines {
		stack := container.NewStack[rune]()

		for _, c := range line {
			if c == '(' || c == '[' || c == '{' || c == '<' {
				stack.Push(c)
			} else {
				v, _ := stack.Pop()

				if *v != openingChars[c] {
					illegalChars = append(illegalChars, c)
				}
			}
		}
	}

	return math.Sum(slice.Map(illegalChars, func(c rune) int { return illegalCharPointMap[c] }))
}

func Part2(input string) int {
	lines := common.GetLines(input)

	totalPoints := make([]int, 0)
	for _, line := range lines {
		stack := container.NewStack[rune]()
		corrupt := false

		for _, c := range line {
			if c == '(' || c == '[' || c == '{' || c == '<' {
				stack.Push(c)
			} else {
				v, _ := stack.Pop()

				if *v != openingChars[c] {
					// ignore corrupt lines
					corrupt = true
					break
				}
			}
		}

		if stack.Len() > 0 && !corrupt {
			points := 0

			for stack.Len() > 0 {
				v, _ := stack.Pop()
				points = points*5 + missingCharPointMap[closingChars[*v]]
			}

			totalPoints = append(totalPoints, points)
		}
	}

	sort.Ints(totalPoints)
	return totalPoints[len(totalPoints)/2]
}
