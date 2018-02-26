package goProbability

import (
	"math/big"
	"testing"
)

func TestFactorial(t *testing.T) {
	a := Factorial(int64(0))
	if a.Cmp(big.NewInt(1)) != 0 {
		t.Errorf("Wrong answer produced: %v, should be: 1", a)
	}
	b := Factorial(int64(1))
	if b.Cmp(big.NewInt(1)) != 0 {
		t.Errorf("Wrong answer produced: %v, should be: 1", a)
	}
	c := Factorial(int64(5))
	if c.Cmp(big.NewInt(5*4*3*2*1)) != 0 {
		t.Errorf("Wrong answer produced: %v, should be: 1", a)
	}

}
