package main

import (
	"fmt"

	goProb "github.com/Miraculin/goProbability"
)

func main() {
	var N, n int64
	var p float64
	fmt.Println("This a trivial example of how the binomial distribution can")
	fmt.Println("approximate a hypergeometric distribution at large N")
	fmt.Println("size of population (N), number of trials? (n), estimated")
	fmt.Println("probability of success per trial(p)")

	_, err := fmt.Scan(&N, &n, &p)
	if err != nil {
		fmt.Println("Error with input", err)
		return
	}
	var binP, hgP float64
	binDist := goProb.Binomial{n, p}
	m := int64(p * float64(N))
	hgDist := goProb.HyperGeom{N, n, m}
	if n > 15 {
		n = 15
	}
	for i := int64(0); i < n; i++ {
		binP, _ = binDist.Prob(i)
		hgP, _ = hgDist.Prob(i)
		fmt.Println("binomial probability: ", binP)
		fmt.Println("hypergeometric probability", hgP)
		fmt.Println("percentage diff.", percentageDifference(binP, hgP))
	}
	// TODO: write example
}

func percentageDifference(x, y float64) float64 {
	ret := (x - y) / ((x + y) / 2.0) * 100
	return ret
}
