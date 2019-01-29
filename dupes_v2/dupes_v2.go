/**
Search suplicate entries in the given argument or file(-s)
Example from the Donovan and Kernighan book
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	references := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, references)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dupes encountered error: %v\n", err)
				continue
			}
			countLines(f, counts, references)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t\t\t%s\n", n, line, references[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, references map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		references[input.Text()] = f.Name()
	}
	//ignore potential errors from input.Err()
}
