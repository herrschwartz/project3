package main

import "fmt"

func main() {
	var in int = 6

	adf := func(a int) int {
		return a + a
	}

	if in > 5 {
		quaduple(adf, in)
	}
	println(adf(in))
}

func quaduple(f func(int) int, a int) {
	fmt.Println(f(a) + f(a))
}
