/*
Assignments example
Example from the Donovan and Kernighan book
*/

package main

import (
	"fmt"
)

//boiling point of water in Fahrentheit degrees
const pointF = 212.0

func main() {
	var f = pointF
	var c = (f - 32) * 5 / 9

	fmt.Printf("Boiling point of water = %g Celcius or %g Fahrentheit\n", c, f)
}
