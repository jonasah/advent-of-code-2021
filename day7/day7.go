package day7

import (
	m "math"
	"strconv"
	"strings"

	"github.com/jonasah/advent-of-code-2021/math"
)

func Part1(input string) int {
	return calculateMinFuelNeeded(input, func(d int) int { return d })
}

func Part2(input string) int {
	return calculateMinFuelNeeded(input, func(d int) int {
		fuelNeeded := 0
		for i := 1; i <= d; i++ {
			fuelNeeded += i
		}
		return fuelNeeded
	})
}

func calculateMinFuelNeeded(input string, fuelFn func(d int) int) int {
	positions := parseInput(input)

	minPos, maxPos := math.MinMax(positions...)
	minFuelNeeded := m.MaxInt

	for p := minPos; p <= maxPos; p++ {
		fuelNeeded := 0
		for _, pos := range positions {
			fuelNeeded += fuelFn(math.Abs(pos - p))
		}

		minFuelNeeded = math.Min(minFuelNeeded, fuelNeeded)
	}

	return minFuelNeeded
}

func parseInput(input string) []int {
	posStrs := strings.Split(input, ",")

	positions := make([]int, len(posStrs))
	for i, posStr := range posStrs {
		pos, _ := strconv.Atoi(posStr)
		positions[i] = pos
	}

	return positions
}
