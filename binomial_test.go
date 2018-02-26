package goProbability

import (
	"math"
	"testing"
)

func TestBinomial(t *testing.T) {
	x := Binomial{10, 0.65}
	if !float64equals(x.Expected(), 6.5) {
		t.Errorf("Wrong answer produced: %v, Expected 6.5", x.Expected())
	}
	a, err := x.Prob(3)
	if err != nil {
		t.Errorf("No error should've been produced")
	}
	if !float64equals(.02120301528515625, a) {
		t.Errorf("Wrong answer produced: %v, Expected 0.0.021203", a)
	}
	b, err := x.Prob(0)
	if !float64equals(b, math.Pow(0.35, float64(10))) {
		t.Errorf("Wrong answer produced: %v, Expected 0.35^10", b)
	}
	if err != nil {
		t.Errorf("No error should've been produced")
	}
	if !float64equals(x.Variance(), 2.275) {
		t.Errorf("Wrong answer produced: %v, Expected:%v", x.Variance(), 10*0.65*0.35)
	}
	_, err = x.Prob(11)
	if err == nil {
		t.Errorf("Error should've been produced, x > n")
	}
}
