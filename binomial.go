package goProbability

import "math"

type Binomial struct {
    n int64
    p float64
}

func (b Binomial) Prob(x int64) (float64, error) {
	if x  > b.n{
	       return 0.0, InvalidInput
	}
    coeff,err := Combination(b.n, x)
	return float64(coeff)*math.Pow(b.p,float64(x))*math.Pow(1-b.p,float64(b.n-x)), err
}

func (b Binomial) Expected() float64 {
	return b.p*float64(b.n)
}

func (b Binomial) Variance() float64 {
	return float64(b.n)*b.p * (1 - b.p)
}
