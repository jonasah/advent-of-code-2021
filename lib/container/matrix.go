package container

import "github.com/jonasah/advent-of-code-2021/lib/iterator"

type MatrixIterValue[T any] struct {
	Point Point
	Value *T
}

type Matrix[T any] interface {
	Rows() int
	Cols() int
	Set(point Point, value T)
	Get(point Point) T
	GetOrDefault(point Point, defaultValue T) T
	Exists(point Point) bool
	Iter() iterator.Iterator[MatrixIterValue[T]]
}

type matrix[T any] struct {
	numRows int
	numCols int
	values  []T
}

func NewMatrix[T any](numRows, numCols int) Matrix[T] {
	return &matrix[T]{numRows, numCols, make([]T, numRows*numCols)}
}

func (m *matrix[T]) Rows() int {
	return m.numRows
}

func (m *matrix[T]) Cols() int {
	return m.numCols
}

func (m *matrix[T]) Set(p Point, value T) {
	m.values[m.index(p)] = value
}

func (m *matrix[T]) SetMapped(p Point, fn func(T) T) {
	m.Set(p, fn(m.Get(p)))
}

func (m *matrix[T]) Get(p Point) T {
	return m.values[m.index(p)]
}

func (m *matrix[T]) GetOrDefault(p Point, def T) T {
	if !m.Exists(p) {
		return def
	}
	return m.Get(p)
}

func (m *matrix[T]) Exists(p Point) bool {
	return p.Row >= 0 && p.Row < m.Rows() && p.Col >= 0 && p.Col < m.Cols()
}

func (m *matrix[T]) Iter() iterator.Iterator[MatrixIterValue[T]] {
	return &matrixIterator[T]{m, 0}
}

func (m *matrix[T]) index(p Point) int {
	return p.Row*m.Cols() + p.Col
}

func (m *matrix[T]) point(index int) Point {
	row := index / m.Cols()
	col := index - row*m.Cols()
	return Point{row, col}
}

type matrixIterator[T any] struct {
	matrix *matrix[T]
	index  int
}

func (it *matrixIterator[T]) HasNext() bool {
	return it.index < len(it.matrix.values)
}

func (it *matrixIterator[T]) Next() *MatrixIterValue[T] {
	point := it.matrix.point(it.index)
	v := &it.matrix.values[it.index]
	it.index++
	return &MatrixIterValue[T]{point, v}
}
