package goProbability

import "math"

const EPSILON float64 = 0.00000001

func float64equals(x, y float64) bool {
	return math.Abs(x-y) < EPSILON
}
