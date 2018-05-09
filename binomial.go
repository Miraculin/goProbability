package goProbability

import "math"

type Binomial struct {
	//n=>0 number of trials
	Bn int64
	//0<=p<=1 probability of success of each trial
	Bp float64
}

//return probability of x successes
func (b Binomial) Prob(x int64) (float64, error) {
	if x > b.Bn {
		return 0.0, InvalidInput
	}
	coeff, err := Combination(b.Bn, x)
	return float64(coeff) * math.Pow(b.Bp, float64(x)) * math.Pow(1-b.Bp, float64(b.Bn-x)), err
}

//return the expected number of successes
func (b Binomial) Expected() float64 {
	return b.Bp * float64(b.Bn)
}

// return the variance in the distribution
func (b Binomial) Variance() float64 {
	return float64(b.Bn) * b.Bp * (1 - b.Bp)
}
