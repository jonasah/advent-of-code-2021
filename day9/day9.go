package day9

import (
	m "math"
	"sort"
	"strconv"
	"strings"

	h "github.com/jonasah/advent-of-code-2021/day9/heightmap"
	"github.com/jonasah/advent-of-code-2021/lib/common"
	"github.com/jonasah/advent-of-code-2021/lib/math"
	"github.com/jonasah/advent-of-code-2021/lib/slice"
)

var topOffset = h.Offset{Row: -1, Col: 0}
var bottomOffset = h.Offset{Row: 1, Col: 0}
var leftOffset = h.Offset{Row: 0, Col: -1}
var rightOffset = h.Offset{Row: 0, Col: 1}
var allOffsets = []h.Offset{topOffset, bottomOffset, leftOffset, rightOffset}

func Part1(input string) int {
	heightMap := parseInput(input)

	lowPoints := getLowPoints(heightMap)

	return math.Sum(slice.Map(lowPoints, func(p h.Point) int { return heightMap.Get(p) + 1 }))
}

func Part2(input string) int {
	heightMap := parseInput(input)

	lowPoints := getLowPoints(heightMap)
	basins := make([]int, len(lowPoints))

	for i, lowPoint := range lowPoints {
		unvisited := make([]h.Point, 1)
		visited := make([]h.Point, 0, 1)

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

func parseInput(input string) h.HeightMap {
	lines := common.GetLines(input)

	heightMap := h.NewHeightMap(len(lines), len(lines[0]))
	for r, line := range lines {
		heightStrings := strings.Split(line, "")

		for c, ht := range heightStrings {
			height, _ := strconv.Atoi(ht)
			heightMap.Set(h.Point{Row: r, Col: c}, height)
		}
	}

	return heightMap
}

func IsLowPoint(heightMap h.HeightMap, p h.Point) bool {
	minNeighborHeight, _ := math.MinMax(slice.Map(allOffsets, func(o h.Offset) int { return heightMap.GetOrDefault(p.Add(o), m.MaxInt) })...)
	return minNeighborHeight > heightMap.Get(p)
}

func getLowPoints(heightMap h.HeightMap) []h.Point {
	points := make([]h.Point, 0)
	for r := 0; r < heightMap.NumRows; r++ {
		for c := 0; c < heightMap.NumCols; c++ {
			p := h.Point{Row: r, Col: c}
			if IsLowPoint(heightMap, p) {
				points = append(points, p)
			}
		}
	}
	return points
}

func shouldVisit(point h.Point, visited []h.Point, unvisited []h.Point, heightMap h.HeightMap) bool {
	exists := heightMap.Exists(point)
	hasVisited := slice.IndexOf(visited, func(p h.Point) bool { return point.Equals(p) }) != -1
	toBeVisited := slice.IndexOf(unvisited, func(p h.Point) bool { return point.Equals(p) }) != -1

	if exists && !hasVisited && !toBeVisited {
		height := heightMap.Get(point)
		return height < 9
	}

	return false
}
