package goProbability

import "math"

const EPSILON float64 = 0.00000001

//returns equality between two float64
func float64equals(x, y float64) bool {
	return math.Abs(x-y) < EPSILON
}

func minInt(x ...int64) int64 {
	min := x[0]
	for _, v := range x {
		if v <= min {
			min = v
		}
	}
	return min
}
