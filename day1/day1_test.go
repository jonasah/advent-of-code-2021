package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jonasah/advent-of-code-2021/common"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 7, Part1(testInput))
	assert.Equal(t, 1709, Part1(realInput))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 5, Part2(testInput))
	assert.Equal(t, 1761, Part2(realInput))
}

var realInput = common.GetInput(1)

const testInput = `199
200
208
210
200
207
240
269
260
263`
