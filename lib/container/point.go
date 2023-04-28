package container

type Point struct {
	Row, Col int
}

type Offset Point

func (p Point) Equals(other Point) bool {
	return p.Row == other.Row && p.Col == other.Col
}

func (p Point) Add(o Offset) Point {
	return Point{Row: p.Row + o.Row, Col: p.Col + o.Col}
}
