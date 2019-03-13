/**
temperature converter. Packages mechanics explained
Example from the Donovan and Kernighan book
*/

package main

import (
	"fmt"

	"./tempconv"
)

func main() {
	fmt.Println(tempconv.FreezingC)
	fmt.Println(tempconv.CtoK(0))
	fmt.Println(tempconv.KtoC(0))
}
