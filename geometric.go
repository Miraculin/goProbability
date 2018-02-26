package goProbability

import "math"

type Geometric struct {
	p float64
}

//Number of trials up to and including the first success version
func (g Geometric) Prob(x int) float64 {
	return math.Pow(1-g.p, float64(x-1)) * g.p
}

func (g Geometric) Expected() float64 {
	return 1 / g.p
}

func (g Geometric) Variance() float64 {
	num := 1 - g.p
	denom := math.Pow(g.p, float64(2))
	return num / denom
}

func (g Geometric) CumulProb(x int) float64 {
	return 1 - math.Pow(1-g.p, float64(x))
}
