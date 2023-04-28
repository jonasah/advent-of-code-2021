package day9

import (
	m "math"
	"sort"
	"strconv"
	"strings"

	"github.com/jonasah/advent-of-code-2021/lib/common"
	"github.com/jonasah/advent-of-code-2021/lib/container"
	"github.com/jonasah/advent-of-code-2021/lib/math"
	"github.com/jonasah/advent-of-code-2021/lib/slice"
)

type heightMap container.Matrix[int]

var topOffset = container.Offset{Row: -1, Col: 0}
var bottomOffset = container.Offset{Row: 1, Col: 0}
var leftOffset = container.Offset{Row: 0, Col: -1}
var rightOffset = container.Offset{Row: 0, Col: 1}
var allOffsets = []container.Offset{topOffset, bottomOffset, leftOffset, rightOffset}

func Part1(input string) int {
	heightMap := parseInput(input)

	lowPoints := getLowPoints(heightMap)

	return math.Sum(slice.Map(lowPoints, func(p container.Point) int { return heightMap.Get(p) + 1 }))
}

func Part2(input string) int {
	heightMap := parseInput(input)

	lowPoints := getLowPoints(heightMap)
	basins := make([]int, len(lowPoints))

	for i, lowPoint := range lowPoints {
		unvisited := make([]container.Point, 1)
		visited := make([]container.Point, 0, 1)

		unvisited[0] = lowPoint

		for len(unvisited) > 0 {
			p := unvisited[0]
			unvisited = unvisited[1:]

			for _, offset := range allOffsets {
				newPos := p.Add(offset)
				if shouldVisit(newPos, visited, unvisited, heightMap) {
					unvisited = append(unvisited, newPos)
				}
			}

			visited = append(visited, p)
		}

		basins[i] = len(visited)
	}

	sort.Ints(basins)
	largestBasins := basins[len(basins)-3:]

	return math.Product(largestBasins)
}

func parseInput(input string) heightMap {
	lines := common.GetLines(input)

	htMap := container.NewMatrix[int](len(lines), len(lines[0]))
	for r, line := range lines {
		heightStrings := strings.Split(line, "")

		for c, ht := range heightStrings {
			height, _ := strconv.Atoi(ht)
			htMap.Set(container.Point{Row: r, Col: c}, height)
		}
	}

	return htMap
}

func IsLowPoint(htMap heightMap, p container.Point) bool {
	minNeighborHeight, _ := math.MinMax(slice.Map(allOffsets, func(o container.Offset) int { return htMap.GetOrDefault(p.Add(o), m.MaxInt) })...)
	return minNeighborHeight > htMap.Get(p)
}

func getLowPoints(htMap heightMap) []container.Point {
	points := make([]container.Point, 0)
	for it := htMap.Iter(); it.HasNext(); {
		p := it.Next().Point
		if IsLowPoint(htMap, p) {
			points = append(points, p)
		}
	}
	return points
}

func shouldVisit(point container.Point, visited []container.Point, unvisited []container.Point, htMap heightMap) bool {
	exists := htMap.Exists(point)
	hasVisited := slice.IndexOf(visited, func(p container.Point) bool { return point.Equals(p) }) != -1
	toBeVisited := slice.IndexOf(unvisited, func(p container.Point) bool { return point.Equals(p) }) != -1

	if exists && !hasVisited && !toBeVisited {
		height := htMap.Get(point)
		return height < 9
	}

	return false
}
