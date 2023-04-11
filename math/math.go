package math

import (
	"math"

	"github.com/jonasah/advent-of-code-2021/lib/slice"
)

func Sum(array []int) int {
	return slice.Reduce(array, func(x, y int) int { return x + y }, 0)
}

func Product(array []int) int {
	return slice.Reduce(array, func(x, y int) int { return x * y }, 1)
}

func MinMax(array ...int) (int, int) {
	minVal, maxVal := math.MaxInt, math.MinInt

	for _, v := range array {
		minVal = Min(minVal, v)
		maxVal = Max(maxVal, v)
	}

	return minVal, maxVal
}

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func Sign(val int) int {
	if val == 0 {
		return 0
	}
	if val < 0 {
		return -1
	}

	return 1
}
