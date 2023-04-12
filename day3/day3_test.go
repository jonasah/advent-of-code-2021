package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jonasah/advent-of-code-2021/lib/common"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 198, Part1(testInput))
	assert.Equal(t, 3429254, Part1(realInput))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 230, Part2(testInput))
	assert.Equal(t, 5410338, Part2(realInput))
}

var realInput = common.GetInput(3)

const testInput = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
