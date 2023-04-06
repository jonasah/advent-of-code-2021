package day8

import (
	"fmt"
	"math"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/jonasah/advent-of-code-2021/lib/common"
)

type segmentCount int

const (
	segmentCount_0_6_9 segmentCount = 6
	segmentCount_1     segmentCount = 2
	segmentCount_2_3_5 segmentCount = 5
	segmentCount_4     segmentCount = 4
	segmentCount_7     segmentCount = 3
	segmentCount_8     segmentCount = 7
)

type wire rune
type wireSet mapset.Set[wire]

type entry struct {
	signalWires    [8][]wireSet // grouped by length
	outputSegments []wireSet
}

func Part1(input string) int {
	entries := parseInput(input)

	occurrences := 0
	for _, entry := range entries {
		for _, o := range entry.outputSegments {
			c := segmentCount(o.Cardinality())

			if c == segmentCount_1 || c == segmentCount_4 || c == segmentCount_7 || c == segmentCount_8 {
				occurrences++
			}
		}
	}

	return occurrences
}

func Part2(input string) int {
	entries := parseInput(input)

	totalOutput := 0
	for _, entry := range entries {
		connections := [10]wireSet{}

		connections[1] = entry.signalWires[segmentCount_1][0]
		topRightCandidates := connections[1].Clone()
		bottomRightCandidates := topRightCandidates.Clone()

		connections[4] = entry.signalWires[segmentCount_4][0]
		topLeftCandidates := connections[4].Difference(connections[1])
		centerCandidates := topLeftCandidates.Clone()

		connections[7] = entry.signalWires[segmentCount_7][0]
		topCandidates := connections[7].Difference(connections[1])

		connections[8] = entry.signalWires[segmentCount_8][0]
		bottomLeftCandidates := connections[8].Difference(connections[1].Union(connections[4]).Union(connections[7]))
		bottomCandidates := bottomLeftCandidates.Clone()

		// ...

		fmt.Println(topCandidates, topLeftCandidates, topRightCandidates, centerCandidates, bottomLeftCandidates, bottomRightCandidates, bottomCandidates)

		connections[0] = topCandidates.Union(topLeftCandidates).Union(bottomLeftCandidates).Union(bottomCandidates).Union(bottomRightCandidates).Union(topRightCandidates)
		connections[2] = topCandidates.Union(topRightCandidates).Union(centerCandidates).Union(bottomLeftCandidates).Union(bottomCandidates)
		connections[3] = topCandidates.Union(topRightCandidates).Union(centerCandidates).Union(bottomRightCandidates).Union(bottomCandidates)
		connections[5] = topCandidates.Union(topLeftCandidates).Union(centerCandidates).Union(bottomRightCandidates).Union(bottomCandidates)
		connections[6] = topCandidates.Union(topLeftCandidates).Union(bottomLeftCandidates).Union(bottomCandidates).Union(bottomRightCandidates).Union(centerCandidates)
		connections[9] = topCandidates.Union(topLeftCandidates).Union(centerCandidates).Union(topRightCandidates).Union(bottomRightCandidates).Union(bottomCandidates)

		fmt.Println(connections)

		outputDigits := make([]int, 0, len(entry.outputSegments))
		for _, o := range entry.outputSegments {
			for n := 0; n < 10; n++ {
				if o.Equal(connections[n]) {
					// fmt.Printf("%v is %v\n", o, n)
					outputDigits = append(outputDigits, n)
					break
				}
			}
		}

		fmt.Println(outputDigits)

		for p, d := range outputDigits {
			totalOutput += d * int(math.Pow10(p))
		}

		break
	}

	return totalOutput
}

func parseInput(input string) []entry {
	lines := common.GetLines(input)

	entries := make([]entry, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " | ")
		signalWires := toWireSets(strings.Split(parts[0], " "))
		outputSegments := toWireSets(strings.Split(parts[1], " "))

		signalWiresMap := [8][]wireSet{}
		for _, ws := range signalWires {
			signalWiresMap[ws.Cardinality()] = append(signalWiresMap[ws.Cardinality()], ws)
		}

		entries[i] = entry{signalWires: signalWiresMap, outputSegments: outputSegments}
	}

	return entries
}

func toWireSets(wireStrings []string) []wireSet {
	wireSets := make([]wireSet, len(wireStrings))
	for i, s := range wireStrings {
		wireSets[i] = mapset.NewSet([]wire(s)...)
	}
	return wireSets
}

func (w wire) String() string {
	return string(w)
}
