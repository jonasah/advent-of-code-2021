package day3

import (
	"sort"
	"strconv"

	"github.com/jonasah/advent-of-code-2021/lib/common"
)

func Part1(input string) int {
	binaryNumbers := common.GetLines(input)

	width := len(binaryNumbers[0])
	bitCount := make([]int, width)

	for _, binary := range binaryNumbers {
		decimal, _ := strconv.ParseInt(binary, 2, 0)

		for d := 0; d < width; d++ {
			value := decimal & (1 << d)

			if value > 0 {
				bitCount[d]++
			} else {
				bitCount[d]--
			}
		}
	}

	gammaRate := 0
	for i, v := range bitCount {
		if v > 0 {
			gammaRate += 1 << i
		}
	}

	epsilonRate := (1 << width) - 1 - gammaRate

	return gammaRate * epsilonRate
}

func Part2(input string) int {
	binaryNumbers := common.GetLines(input)

	width := len(binaryNumbers[0])
	sort.Strings(binaryNumbers)

	mostCommonCandidates := binaryNumbers
	leastCommonCandidates := binaryNumbers
	for d := 0; d < width; d++ {
		if len(mostCommonCandidates) > 1 {
			mostCommonCandidates = findMostCommonCandidates(mostCommonCandidates, d)
		}
		if len(leastCommonCandidates) > 1 {
			leastCommonCandidates = findLeastCommonCandidates(leastCommonCandidates, d)
		}
	}

	oxygenRating, _ := strconv.ParseInt(mostCommonCandidates[0], 2, 0)
	co2Rating, _ := strconv.ParseInt(leastCommonCandidates[0], 2, 0)

	return int(oxygenRating) * int(co2Rating)
}

func findMostCommonCandidates(binaryNumbers []string, d int) []string {
	index := sort.Search(len(binaryNumbers), func(i int) bool {
		return binaryNumbers[i][d] == '1'
	})

	if index > len(binaryNumbers)/2 {
		return binaryNumbers[:index]
	}

	return binaryNumbers[index:]
}

func findLeastCommonCandidates(binaryNumbers []string, d int) []string {
	index := sort.Search(len(binaryNumbers), func(i int) bool {
		return binaryNumbers[i][d] == '1'
	})

	if index > len(binaryNumbers)/2 {
		return binaryNumbers[index:]
	}

	return binaryNumbers[:index]
}
