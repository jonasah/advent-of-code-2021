package day5

type point struct {
	x, y int
}

type line struct {
	p0, p1 point
}

func (l line) isHorizontalOrVertical() bool {
	return l.p0.y == l.p1.y || l.p0.x == l.p1.x
}
