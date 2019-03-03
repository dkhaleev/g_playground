/**
understanding pointers
Example from the Donovan and Kernighan book

*/

package main

import (
	"fmt"
)

func main() {
	x := 1
	p := &x
	fmt.Println(*p)

	*p = 2
	fmt.Println(x)

	fmt.Println(incrP(&x))
}

// function increases any given variable indirectly via pointer
func incrP(p *int) int {
	*p++
	return *p
}
