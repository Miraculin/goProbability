package goProbability

type HyperGeom struct {
	N int64
	n int64
	m int64
}

//return the probability that x successes occur
func (h HyperGeom) Prob(x int64) (float64, error) {
	low := minInt(0, h.n+h.m-h.N)
	high := minInt(h.m, h.n)
	if x > high || x < low {
		return 0.0, InvalidInput
	}
	coeff1, err := Combination(h.m, x)
	coeff2, err := Combination(h.N-h.m, h.n-x)
	coeff3, err := Combination(h.N, h.n)
	return float64(coeff1) * float64(coeff2) / float64(coeff3), err
}

//return the expected number of successes
func (h HyperGeom) Expected() float64 {
	return float64(h.m) * float64(h.n) / float64(h.N)
}

//return the variance of the distribution
func (h HyperGeom) Variance() float64 {
	return h.Expected() * (1 - float64(h.m)/float64(h.N)) * (float64(h.N-h.n) / float64(h.N-1))
}
