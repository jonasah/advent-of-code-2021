package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jonasah/advent-of-code-2021/lib/common"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 5934, Part1(testInput))
	assert.Equal(t, 346063, Part1(realInput))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 26984457539, Part2(testInput))
	assert.Equal(t, 1572358335990, Part2(realInput))
}

var realInput = common.GetInput(6)

const testInput = `3,4,3,1,2`
