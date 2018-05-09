package goProbability

import "math"

type Poisson struct {
	Lambda float64
}

//return the probability that x trials occur
func (p Poisson) Prob(x int64) (float64, error) {
	if x < 0 {
		return 0.0, InvalidInput
	}
	num := math.Pow(p.Lambda, float64(x)) * math.Exp(-p.Lambda)
	denom := Factorial(x).Int64()
	return num / float64(denom), nil
}

//return the expected number of successes
func (p Poisson) Expected() float64 {
	return p.Lambda
}

//return the variance of the distribution
func (p Poisson) Variance() float64 {
	return p.Lambda
}
