package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	z -= (z*z - x) / (2 * z)
	z -= (z*z - x) / (2 * z)
	z -= (z*z - x) / (2 * z)
	z -= (z*z - x) / (2 * z)
	z -= (z*z - x) / (2 * z)
	z -= (z*z - x) / (2 * z)
	return z
}

func execLoopsFunc() {
	fmt.Println(Sqrt(3))
}
