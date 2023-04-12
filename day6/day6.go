package day6

import (
	"strconv"
	"strings"

	"github.com/jonasah/advent-of-code-2021/lib/math"
)

func Part1(input string) int {
	return doIt(input, 80)
}

func Part2(input string) int {
	return doIt(input, 256)
}

func doIt(input string, numDays int) int {
	buckets := parseInput(input)

	for day := 0; day < numDays; day++ {
		numNew := buckets[0]

		for i := 1; i < 9; i++ {
			buckets[i-1] = buckets[i]
		}

		buckets[6] += numNew
		buckets[8] = numNew
	}

	return math.Sum(buckets[:])
}

func parseInput(input string) [9]int {
	timerStrs := strings.Split(input, ",")

	buckets := [9]int{}
	for _, timerStr := range timerStrs {
		t, _ := strconv.Atoi(timerStr)
		buckets[t]++
	}

	return buckets
}
