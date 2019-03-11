/**
Linux-like echo v.4 command
Example from the Donovan and Kernighan book

*/

package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "New line character")
var s = flag.String("s", " ", "Separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *s))

	if !*n {
		fmt.Println()
	}
}
