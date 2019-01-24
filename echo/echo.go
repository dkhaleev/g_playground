/*Linux-like echo command
Example from the Donovan and Kernighan book
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, separator string

	s = os.Args[0] + " : "
	for i := 1; i < len(os.Args); i++ {
		s += separator + os.Args[i]
		separator = " "
	}

	fmt.Println(s)
}
