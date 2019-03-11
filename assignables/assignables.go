/*
Assignables and variables
Example from the Donovan and Kernighan book
*/

package main

import (
	"fmt"
)

func main() {

	metals := []string{"brass", "copper", "silver", "thungsten"}

	for j := range metals {
		metal := &metals[j]
		fmt.Println(*metal)

	}

}
