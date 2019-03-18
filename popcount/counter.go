/**
count of all set bits in byte. Understanding init()
Example from the Donovan and Kernighan book
*/

package main

import (
	"fmt"
	"os"
	"strconv"

	"./popcount"
)

func main() {
	for _, arg := range os.Args[1:] {
		i, err := strconv.ParseUint(arg, 10, 32)
		if err != nil {
			fmt.Println("Argument parse error")
			os.Exit(1)
		}

		fmt.Printf("set bits total in byte %d\n", popcount.Popcount(i))
	}
}
