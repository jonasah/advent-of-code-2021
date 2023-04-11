package heightmap

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

type HeightMap struct {
	NumRows int
	NumCols int
	heights []int
}

func NewHeightMap(numRows, numCols int) HeightMap {
	return HeightMap{numRows, numCols, make([]int, numRows*numCols)}
}

func (h *HeightMap) Set(p Point, height int) {
	h.heights[h.index(p)] = height
}

func (h *HeightMap) Get(p Point) int {
	return h.heights[h.index(p)]
}

func (h *HeightMap) GetOrDefault(p Point, def int) int {
	if !h.Exists(p) {
		return def
	}
	return h.Get(p)
}

func (h *HeightMap) Exists(p Point) bool {
	return p.Row >= 0 && p.Row < h.NumRows && p.Col >= 0 && p.Col < h.NumCols
}

func (h *HeightMap) index(p Point) int {
	return p.Row*h.NumCols + p.Col
}
