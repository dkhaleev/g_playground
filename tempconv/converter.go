/**
temperature converter. Packages mechanics explained
Example from the Donovan and Kernighan book
*/

package main

import (
	"fmt"
	"os"
	"strconv"

	"./tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		fmt.Printf("%s = %s, %s = %s", f, tempconv.FtoC(f), c, tempconv.CtoF(c))
	}
}
