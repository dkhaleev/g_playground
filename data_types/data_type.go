/**
Fundametal data types
Example from the Donovan and Kernighan book
*/
package main

import (
	"fmt"
)

var a uint8 = 1<<1 | 1<<5
var b uint8 = 1<<1 | 1<<2

func main() {
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", a&b)
	fmt.Printf("%08b\n", a|b)
	fmt.Printf("%08b\n", a^b)
	fmt.Printf("%08b\n", a&^b)
}
