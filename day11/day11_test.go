package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jonasah/advent-of-code-2021/lib/common"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 1656, Part1(testInput))
	// assert.Equal(t, -1, Part1(realInput))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, -1, Part2(testInput))
	// assert.Equal(t, -1, Part2(realInput))
}

var realInput = common.GetInput(11)

const testInput = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`
