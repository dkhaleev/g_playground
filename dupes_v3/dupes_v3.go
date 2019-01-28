/**
Search suplicate entries in the given argument or file(-s)
Read lines
Example from the Donovan and Kernighan book
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, error := ioutil.ReadFile(filename)
		if error != nil {
			fmt.Fprintf(os.Stderr, "Dupes encountered an error %v \n", error)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
