package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jonasah/advent-of-code-2021/lib/common"
	"github.com/jonasah/advent-of-code-2021/lib/container"
	"github.com/jonasah/advent-of-code-2021/lib/slice"
)

type grid container.Matrix[int]

const maxEnergy = 9

var adjOffsets = []container.Offset{
	{Row: -1, Col: -1},
	{Row: -1, Col: 0},
	{Row: -1, Col: 1},
	{Row: 0, Col: -1},
	{Row: 0, Col: 1},
	{Row: 1, Col: -1},
	{Row: 1, Col: 0},
	{Row: 1, Col: 1}}

func Part1(input string) int {
	grid := parseInput(input)

	numFlashes := 0
	for steps := 0; steps < 100; steps++ {
		incLevels(grid)
		numFlashes += flash(grid)
		resetLevels(grid)

		break
	}

	fmt.Println(grid, numFlashes)

	return numFlashes
}

func Part2(input string) int {
	return 0
}

func parseInput(input string) grid {
	lines := common.GetLines(input)

	grid := container.NewMatrix[int](len(lines), len(lines[0]))
	for r, line := range lines {
		heightStrings := strings.Split(line, "")

		for c, ht := range heightStrings {
			height, _ := strconv.Atoi(ht)
			grid.Set(container.Point{Row: r, Col: c}, height)
		}
	}

	return grid
}

func incLevels(g grid) {
	for it := g.Iter(); it.HasNext(); {
		*it.Next().Value++
	}
}

func flash(g grid) int {
	numFlashes := 0
	for it := g.Iter(); it.HasNext(); {
		itt := it.Next()

		if *itt.Value == maxEnergy {
			// numFlashes += len(adj)
			numFlashes++

			// adj := getAdjacentPoints(g, itt.Point)
		}
	}
	return numFlashes
}

func resetLevels(g grid) {
	for it := g.Iter(); it.HasNext(); {
		v := it.Next().Value
		if *v > maxEnergy {
			*v = 0
		}
	}
}

func getAdjacentPoints(g grid, p container.Point) []container.Point {
	return slice.Filter(
		slice.Map(adjOffsets, func(o container.Offset) container.Point { return p.Add(o) }),
		func(p container.Point) bool { return g.Exists(p) })
}
