package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		// if this 'if' condition is removed, we get the following error:
		// panic: runtime error: invalid memory address or nil pointer dereference
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I
	fmt.Println(i) // prints <nil>

	var t *T
	fmt.Println(t) // prints <nil>
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%+v, %T)\n", i, i)
}
