/**
Linux-like echo v.2 command
Example from the Donovan and Kernighan book

*/

package main

import (
	"fmt"
	"os"
)

func main() {
	str, sep := "", ""

	for _, arg := range os.Args[1:] {
		str += sep + arg
	}

	fmt.Printf("Input argument is: %s", str)
}
