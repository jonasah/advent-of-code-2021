package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jonasah/advent-of-code-2021/common"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 5, Part1(testInput))
	assert.Equal(t, 6841, Part1(realInput))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 12, Part2(testInput))
	assert.Equal(t, 19258, Part2(realInput))
}

var realInput = common.GetInput(5)

const testInput = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
