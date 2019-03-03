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
}
