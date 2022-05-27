package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	zsq := z
	for i := 0; (i < 10) && ((math.Abs(zsq-x) / x) > 0.000000001); i++ {
		z -= ((z * z) - x) / (2 * z)
		zsq = z * z
		fmt.Printf("%d : %.19f - sqr = %.19f [ zsq/x = %.19f ]\n", i, z, (z * z), (zsq / x))
	}
	return (z)
}

func main() {
	fmt.Printf("Our function result : %.9f\n", Sqrt(2))
	fmt.Printf("Math func result    : %.9f\n", math.Sqrt(2))
}
