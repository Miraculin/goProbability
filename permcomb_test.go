package goProbability

import "testing"

func TestPermcomb(t *testing.T) {
	n1 := int64(10)
	r1 := int64(2)
	r2 := int64(11)
	r3 := []int64{2, 2, 5}
	r4 := []int64{2, 4, 5}
	a, err := Partition(n1, r1)
	if err != nil {
		t.Errorf("No errors should've occured")
	}
	v := int64(5 * 9)
	if a != v {
		t.Errorf("Wrong answer produced, should be:%v", v)
	}
	_, err = Partition(n1, r2)
	if err == nil {
		t.Errorf("Error should've occured, r>n")
	}
	b, err := Partition(n1, r3...)
	if err != nil {
		t.Errorf("No errors should've occured")
	}
	if b != 7560 {
		t.Errorf("Wrong answer produced, should be: 7560")
	}

	_, err = Partition(n1, r4...)
	if err == nil {
		t.Errorf("Error should've been produced, sum of r4>n")
	}
	c, err := Permutation(n1, r1)
	if err != nil {
		t.Errorf("No errors should've occured")
	}
	if c != int64(10*9) {
		t.Errorf("Wrong answer produced, should be:%v", 90)
	}
	_, err = Permutation(n1, r2)
	if err == nil {
		t.Errorf("Error should've occured, r>n")
	}
	d, err := Combination(n1, r1)
	if err != nil {
		t.Errorf("No errors should've occured")
	}
	if d != v {
		t.Errorf("Wrong answer produced, should be:%v", v)
	}
	_, err = Combination(n1, r2)
	if err == nil {
		t.Errorf("Error should've been produced, r2>n")
	}
}
