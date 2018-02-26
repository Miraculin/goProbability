package goProbability

func Permutation(n, r int64) (int64, error) {
	if r > n {
		return 0, InvalidInput
	}
	num := Factorial(n)
	denom := Factorial(n - r)
	return (num.Div(num, denom)).Int64(), nil
}

func Combination(n, r int64) (int64, error) {
	return Partition(n, r)
}

func Partition(n int64, r ...int64) (int64, error) {
	rsum := int64(0)
	for _, nk := range r {
		rsum += nk
	}
	if rsum > n {
		return 0, InvalidInput
	}
	num := Factorial(n)
	for _, nk := range r {
		denom := Factorial(nk)
		num.Div(num, denom)
	}
	num.Div(num, Factorial(n-rsum))
	return num.Int64(), nil
}
