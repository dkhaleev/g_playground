/*Linux-like echo command
Example from the Donovan and Kernighan book
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	var separator string
	// 2nd homework
	// var s, separator string

	//1st homework
	// s = os.Args[0] + " : "
	for i := 1; i < len(os.Args); i++ {
		// s += separator + os.Args[i]
		separator = " "
		fmt.Printf("index #%d%svalue:%s%s \n", i, separator, separator, os.Args[i])
	}

	// 2nd homework
	// fmt.Println(s)
}
