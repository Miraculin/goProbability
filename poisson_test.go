package goProbability

import (
	"testing"
)

func TestPoisson(t *testing.T) {
	x := Poisson{10}
	if !float64equals(x.Expected(), 10) {
		t.Errorf("Wrong answer produced: %v, Expected 40", x.Expected())
	}
	_, err := x.Prob(-1)
	if err == nil {
		t.Errorf("error should've been produced")
	}
	a, err := x.Prob(5)
	if err != nil {
		t.Errorf("No error should've been produced")
	}
	if !float64equals(a, 0.03783327480) {
		t.Errorf("Wrong answer produced: %v, Expected 0.03783327480", a)
	}
	if !float64equals(x.Variance(), 10) {
		t.Errorf("Wrong answer produced: %v, Expected:%v", x.Variance(), 10)
	}
}
