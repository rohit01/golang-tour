package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["answer"] = 42
	fmt.Println("The value:", m["answer"])

	m["answer"] = 380
	fmt.Println("The value:", m["answer"])

	delete(m, "answer")
	fmt.Println("The value:", m["answer"])

	v, ok := m["answer"]
	fmt.Println("The value:", v, "Present?", ok)

	v = m["r"]
	fmt.Println("The value:", v)

	x, y, z := r3()
	fmt.Println(x, y, z)
}

func r3() (int, int, int) {
	return 1, 2, 3
}
