package goProbability

import "math"

type NegBinom struct {
	//0<p<1 probability of success of each trial
	Np float64
	//r>0 number of success until experiment stops
	Nr int64
}

//return the probability that x trials occur
func (n NegBinom) Prob(x int64) (float64, error) {
	if x < n.Nr {
		return 0.0, InvalidInput
	}
	coeff, err := Combination(x-1, n.Nr-1)
	return float64(coeff) * math.Pow(n.Np, float64(n.Nr)) * math.Pow(1-n.Np, float64(x-n.Nr)), err
}

//return the expected number of successes
func (n NegBinom) Expected() float64 {
	return float64(n.Nr) / n.Np
}

//return the variance of the distribution
func (n NegBinom) Variance() float64 {
	return float64(n.Nr) * (1 - n.Np) / (math.Pow(n.Np, float64(2)))
}
