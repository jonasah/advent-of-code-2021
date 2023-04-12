package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jonasah/advent-of-code-2021/lib/common"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 150, Part1(testInput))
	assert.Equal(t, 1893605, Part1(realInput))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 900, Part2(testInput))
	assert.Equal(t, 2120734350, Part2(realInput))
}

var realInput = common.GetInput(2)

const testInput = `forward 5
down 5
forward 8
up 3
down 8
forward 2`
