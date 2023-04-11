package day9

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jonasah/advent-of-code-2021/common"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 15, Part1(testInput))
	assert.Equal(t, 498, Part1(realInput))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 1134, Part2(testInput))
	assert.Equal(t, 1071000, Part2(realInput))
}

var realInput = common.GetInput(9)

const testInput = `2199943210
3987894921
9856789892
8767896789
9899965678`
