package goProbability

type NegBinom struct {
	//0<p<1 probability of success of each trial
	p float64
	//r>0 number of failures until experiment stops
	r int
}

//return the probability that x successes appear
func (n NegBinom) Prob(x int) float64 {
	return 0.0
}

//return the expected number of successes
func (n NegBinom) Expected() float64 {
	return 0.0
}

//return the variance of the distribution
func (n NegBinom) Variance() float64 {
	return 0.0
}
