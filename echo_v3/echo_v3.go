/**
Linux-like echo v.3 command
Example from the Donovan and Kernighan book

*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("input argument is %s", strings.Join(os.Args[1:], " "))
}
