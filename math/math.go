package math

import "math"

func Sum(array []int) int {
	sum := 0
	for _, v := range array {
		sum += v
	}
	return sum
}

func MinMax(array []int) (int, int) {
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
