package goProbability

import (
	"testing"
)

func TestBernoulli(t *testing.T) {
	x := Bernoulli{0.65}
	if x.Expected() != 0.65 {
		t.Errorf("Wrong answer produced: %v, Expected 0.65", x.Expected())
	}
	a, err := x.Prob(1)
	if err != nil {
		t.Errorf("No error should've been produced")
	}
	if a != 0.65 {
		t.Errorf("Wrong answer produced: %v, Expected 0.65", a)
	}
	b, err := x.Prob(0)
	if b != 0.35 {
		t.Errorf("Wrong answer produced: %v, Expected 0.35", b)
	}
	if err != nil {
		t.Errorf("No error should've been produced")
	}
	if !float64equals(x.Variance(), 0.65*0.35) {
		t.Errorf("Wrong answer produced: %v, Expected:%v", x.Variance(), 0.65*0.35)
	}
	_, err = x.Prob(2)
	if err == nil {
		t.Errorf("Error should've been produced, x not 0 or 1")
	}
}
