package goProbability

import "math"

type Bernoulli struct {
	// 0<=p<=1, success probability
	Bp float64
}

//returns probability of a Bernoulli trial or an error if the input's invalid
func (b Bernoulli) Prob(x int) (float64, error) {
	if x == 0 || x == 1 {
		return math.Pow(b.Bp, float64(x)) * (math.Pow(1-b.Bp, float64(1-x))), nil
	}
	return 0.0, InvalidInput
}

//returns the expected value of a Bernoulli trial
func (b Bernoulli) Expected() float64 {
	return b.Bp
}

// returns the variance of a Bernoulli trial
func (b Bernoulli) Variance() float64 {
	return b.Bp * (1 - b.Bp)
}
