/**
main file for types example. Temperature conversion
Example from the Donovan and Kernighan book

*/

package main

import (
	"fmt"

	"./conv"
)

func main() {
	boilingF := conv.CtoF(conv.BoilingC)
	fmt.Printf("%g\n", boilingF)

	c := conv.FtoC(212.0)
	fmt.Println(c.String())
}
