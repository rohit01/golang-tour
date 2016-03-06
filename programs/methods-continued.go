package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

/*
 * // Not allowed. New methods are allowed only on global types.
 * func (i int) Absolute() int {
 * 	if i < 0 {
 * 		return -i
 * 	}
 * 	return i
 * }
 */

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
