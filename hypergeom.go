package goProbability

type HyperGeom struct {
	HN int64
	Hn int64
	Hm int64
}

//return the probability that x successes occur
func (h HyperGeom) Prob(x int64) (float64, error) {
	low := minInt(0, h.Hn+h.Hm-h.HN)
	high := minInt(h.Hm, h.Hn)
	if x > high || x < low {
		return 0.0, InvalidInput
	}
	coeff1, err := Combination(h.Hm, x)
	coeff2, err := Combination(h.HN-h.Hm, h.Hn-x)
	coeff3, err := Combination(h.HN, h.Hn)
	return float64(coeff1) * float64(coeff2) / float64(coeff3), err
}

//return the expected number of successes
func (h HyperGeom) Expected() float64 {
	return float64(h.Hm) * float64(h.Hn) / float64(h.HN)
}

//return the variance of the distribution
func (h HyperGeom) Variance() float64 {
	return h.Expected() * (1 - float64(h.Hm)/float64(h.HN)) * (float64(h.HN-h.Hn) / float64(h.HN-1))
}
