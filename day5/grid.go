package day5

import "github.com/jonasah/advent-of-code-2021/lib/math"

type grid struct {
	points [1000][1000]int
}

func (g *grid) addLine(line line) {
	xdir := math.Sign(line.p1.x - line.p0.x)
	ydir := math.Sign(line.p1.y - line.p0.y)

	for p := line.p0; p.x != line.p1.x || p.y != line.p1.y; {
		g.points[p.y][p.x]++

		p.x += xdir
		p.y += ydir
	}

	g.points[line.p1.y][line.p1.x]++
}

func (g *grid) getNumPointsWithOverlap() int {
	points := 0
	for _, row := range g.points {
		for _, column := range row {
			if column > 1 {
				points++
			}
		}
	}
	return points
}
