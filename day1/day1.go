package day1

import (
	"strconv"

	"github.com/jonasah/advent-of-code-2021/common"
)

func Part1(input string) int {
	return compareDepth(parseDepths(common.GetLines(input)))
}

func Part2(input string) int {
	depths := parseDepths(common.GetLines(input))

	combinedDepths := make([]int, len(depths)-2)
	for i, currDepth := range depths {
		if i > 1 {
			prevPrevDepth := depths[i-2]
			prevDepth := depths[i-1]
			combinedDepths[i-2] = prevPrevDepth + prevDepth + currDepth
		}
	}

	return compareDepth(combinedDepths)
}

func parseDepths(lines []string) []int {
	depths := make([]int, len(lines))
	for i, line := range lines {
		val, _ := strconv.Atoi(line)
		depths[i] = val
	}
	return depths
}

func compareDepth(depths []int) int {
	depthInc := 0
	for i, currDepth := range depths {
		if i > 0 {
			prevDepth := depths[i-1]

			if currDepth > prevDepth {
				depthInc++
			}
		}
	}
	return depthInc
}
