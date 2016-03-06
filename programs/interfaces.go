package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	fmt.Printf("%+v, %+v\n", v, v.Abs())
	p := &v
	fmt.Printf("%+v, %+v\n", p, p.Abs())

	a = f
	fmt.Printf("a = f: %+v, %+v\n", a, a.Abs())
	a = &f
	fmt.Printf("a = &f: %+v, %+v\n", a, a.Abs())
	// ./interfaces.go:25: cannot use v (type Vertex) as type Abser in assignment:
	//	Vertex does not implement Abser (Abs method has pointer receiver)
	//	a = v
	//	fmt.Printf("a = v: %+v, %+v\n", a, a.Abs())
	a = &v
	fmt.Printf("a = &v: %+v, %+v\n", a, a.Abs())

}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
