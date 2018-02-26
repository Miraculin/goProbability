package goProbability

import (
	"math"
	"testing"
)

func TestGeometric(t *testing.T) {
	x := Geometric{0.36}
	if !float64equals(x.Expected(), 1/0.36) {
		t.Errorf("Wrong answer produced: %v, Expected 2.77777777", x.Expected())
	}
	a := x.Prob(3)
	if !float64equals(0.36*math.Pow(0.64, 2), a) {
		t.Errorf("Wrong answer produced: %v, Expected 0.09437184", a)
	}
	b := x.Prob(1)
	if !float64equals(b, 0.36*math.Pow(0.64, 0)) {
		t.Errorf("Wrong answer produced: %v, Expected 0.2304", b)
	}

	if !float64equals(x.Variance(), (1-0.36)/(0.36*0.36)) {
		t.Errorf("Wrong answer produced: %v, Expected:%v", x.Variance(), 4.93827)
	}
	c := x.CumulProb(3)
	if !float64equals(c, 1-math.Pow(1-0.36, float64(3))) {
		t.Errorf("Wrong answer produced: %v, Expected:%v", c, math.Pow(1-0.36, float64(4)))
	}
}
