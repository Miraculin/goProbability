package goProbability

import "math/big"

func Factorial(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	return big.NewInt(n).Mul(big.NewInt(n), Factorial(n-1))
}
