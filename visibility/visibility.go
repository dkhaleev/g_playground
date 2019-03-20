/**
understanding visibility
Example from the Donovan and Kernighan book
*/

package main

import (
	"fmt"
	"log"
	"os"
)

var cwd = "test cwd"

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error: Getwd: %v", err)
	}

	log.Printf("Current working dir: %s", cwd)
}

func main() {

	x := &cwd
	for i := 0; i < len(*x); i++ {
		x := (*x)[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
}
