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
var metals = []string{"silver", "copper", "steel", "brass", "zinc"}

func main() {
	o := 0666
	var p = int64(0xdeadbeef)
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", a&b)
	fmt.Printf("%08b\n", a|b)
	fmt.Printf("%08b\n", a^b)
	fmt.Printf("%08b\n", a&^b)

	for i := uint8(0); i < 8; i++ {
		if a&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	for k := len(metals) - 1; k > 0; k-- {
		fmt.Println(metals[k])
	}

	fmt.Printf("%d %[1]o %#[1]o\n", o)
	fmt.Printf("%d %[1]o %#[1]o %#[1]x %#[1]X %08[1]b\n", p)
}
