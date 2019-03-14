/**
inches to centimeters converter. Packages mechanics explained
Example from the Donovan and Kernighan book
*/

package main

import (
	"fmt"
	"os"
	"strconv"

	"./in2cmconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "converter: %v\n", err)
			os.Exit(1)
		}
		i := in2cmconv.Inch(t)
		cm := in2cmconv.Centimeter(t)

		fmt.Printf("%s = %s, %s = %s", i, in2cmconv.In2cm(i), cm, in2cmconv.Cm2in(cm))
	}
}
