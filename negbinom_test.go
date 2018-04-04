package goProbability

import (
	"math"
	"testing"
)

func TestNegBinom(t *testing.T) {
	x := NegBinom{0.36, 5}
	if !float64equals(x.Expected(), 5/0.36) {
		t.Errorf("Wrong answer produced: %v, Expected 2.77777777", x.Expected())
	}
	_, err := x.Prob(3)
	if err == nil {
		t.Errorf("error should've been produced")
	}
	a, err := x.Prob(7)
	if err != nil {
		t.Errorf("No error should've been produced")
	}
	if !float64equals(a, 15*math.Pow(0.36, float64(5))*math.Pow(0.64, float64(2))) {
		t.Errorf("Wrong answer produced: %v, Expected 0.0371504185", a)
	}
	if !float64equals(x.Variance(), 5*(1-0.36)/(0.36*0.36)) {
		t.Errorf("Wrong answer produced: %v, Expected:%v", x.Variance(), 24.69135)
	}
}
