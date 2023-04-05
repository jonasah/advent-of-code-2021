package day5

import (
	"regexp"
	"strconv"

	"github.com/jonasah/advent-of-code-2021/common"
)

func Part1(input string) int {
	return doIt(input, func(l line) bool { return l.isHorizontalOrVertical() })
}

func Part2(input string) int {
	return doIt(input, func(line) bool { return true })
}

func doIt(input string, lineFilter func(line) bool) int {
	lines := parseInput(input)

	grid := &grid{}
	for _, line := range lines {
		if lineFilter(line) {
			grid.addLine(line)
		}
	}

	return grid.getNumPointsWithOverlap()
}

func parseInput(input string) []line {
	inputLines := common.GetLines(input)

	regex := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

	lines := make([]line, len(inputLines))
	for i, inputLine := range inputLines {
		matches := regex.FindStringSubmatch(inputLine)
		x0, _ := strconv.Atoi(matches[1])
		y0, _ := strconv.Atoi(matches[2])
		x1, _ := strconv.Atoi(matches[3])
		y1, _ := strconv.Atoi(matches[4])
		lines[i] = line{point{x0, y0}, point{x1, y1}}
	}

	return lines
}
