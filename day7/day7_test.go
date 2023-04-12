package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jonasah/advent-of-code-2021/lib/common"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 37, Part1(testInput))
	assert.Equal(t, 345197, Part1(realInput))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 168, Part2(testInput))
	assert.Equal(t, 96361606, Part2(realInput))
}

var realInput = common.GetInput(7)

const testInput = `16,1,2,0,4,2,7,1,2,14`
