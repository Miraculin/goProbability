package goProbability

import "math"

type NegBinom struct {
	//0<p<1 probability of success of each trial
	p float64
	//r>0 number of success until experiment stops
	r int64
}

//return the probability that x trials occur
func (n NegBinom) Prob(x int64) (float64, error) {
	if x < n.r {
		return 0.0, InvalidInput
	}
	coeff, err := Combination(x-1, n.r-1)
	return float64(coeff) * math.Pow(n.p, float64(n.r)) * math.Pow(1-n.p, float64(x-n.r)), err
}

//return the expected number of successes
func (n NegBinom) Expected() float64 {
	return float64(n.r) / n.p
}

//return the variance of the distribution
func (n NegBinom) Variance() float64 {
	return float64(n.r) * (1 - n.p) / (math.Pow(n.p, float64(2)))
}
