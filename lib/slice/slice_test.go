package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jonasah/advent-of-code-2021/lib/math"
	"github.com/jonasah/advent-of-code-2021/lib/slice"
)

func TestReduce(t *testing.T) {
	a := []int{1, 6, 9, 3, 4}
	assert.Equal(t, 9, slice.Reduce(a, math.Max, 0))
}

func TestMap(t *testing.T) {
	a := []int{1, 6, 9, 3, 4}
	m := slice.Map(a, func(val int) int { return val * val })
	assert.ElementsMatch(t, []int{1, 36, 81, 9, 16}, m)
}

func TestIndexOf(t *testing.T) {
	a := []int{1, 6, 9, 3, 4}

	i := slice.IndexOf(a, func(val int) bool { return val == 9 })
	assert.Equal(t, 2, i)

	i = slice.IndexOf(a, func(val int) bool { return val == 7 })
	assert.Equal(t, -1, i)
}

func TestFilter(t *testing.T) {
	a := []int{1, 6, 9, 3, 4}
	f := slice.Filter(a, func(val int) bool { return val > 5 })
	assert.ElementsMatch(t, []int{6, 9}, f)
}
