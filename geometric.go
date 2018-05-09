package goProbability

import "math"

type Geometric struct {
	//0<p<1 probability of success
	Gp float64
}

//Number of trials up to and including the first success version
//return probability that the x-th trial is the first success
func (g Geometric) Prob(x int) float64 {
	return math.Pow(1-g.Gp, float64(x-1)) * g.Gp
}

//return expected number of trials until the first success
func (g Geometric) Expected() float64 {
	return 1 / g.Gp
}

//return the variance in the distribution
func (g Geometric) Variance() float64 {
	num := 1 - g.Gp
	denom := math.Pow(g.Gp, float64(2))
	return num / denom
}

//return the cumulative probability X<=x
func (g Geometric) CumulProb(x int) float64 {
	return 1 - math.Pow(1-g.Gp, float64(x))
}
