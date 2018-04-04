package goProbability

import (
	"testing"
)

func TestHyperGeom(t *testing.T) {
	x := HyperGeom{50, 30, 10}
	if !float64equals(x.Expected(), 6) {
		t.Errorf("Wrong answer produced: %v, Expected 6", x.Expected())
	}
	_, err := x.Prob(60)
	if err == nil {
		t.Errorf("error should've been produced")
	}
	a, err := x.Prob(5)
	if err != nil {
		t.Errorf("No error should've been produced")
	}
	if !float64equals(a, 0.215085007184925) {
		t.Errorf("Wrong answer produced: %v, Expected 0.215085007184925", a)
	}
	if !float64equals(x.Variance(), 96.0/49) {
		t.Errorf("Wrong answer produced: %v, Expected:%v", x.Variance(), 96.0/49)
	}
	y := HyperGeom{50, 10, 30}
	if !float64equals(y.Expected(), 6) {
		t.Errorf("Wrong answer produced:%v, Expected:6", y.Expected())
	}
	_, err = y.Prob(30)
	if err == nil {
		t.Errorf("error should've been produced")
	}
	b, err := y.Prob(5)
	if err != nil {
		t.Errorf("No error should've been produced")
	}
	if !float64equals(b, 0.215085007184925) {
		t.Errorf("Wrong answer produced: %v, Expected 0.215085007184925", a)
	}
	if !float64equals(y.Variance(), 96.0/49) {
		t.Errorf("Wrong answer produced: %v, Expected:%v", x.Variance(), 96.0/49)
	}
}
