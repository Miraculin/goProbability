package goProbability

import "math"

type Bernoulli struct {
	p float64
}

func (b Bernoulli) Prob(x int) (float64, error) {
	if x == 0 || x == 1 {
		return math.Pow(b.p, float64(x)) * (math.Pow(1-b.p, float64(1-x))), nil
	}
	return 0.0, InvalidInput
}

func (b Bernoulli) Expected() float64 {
	return b.p
}

func (b Bernoulli) Variance() float64 {
	return b.p * (1 - b.p)
}
