/*boolean playground
Chapter 3.4 Example from the Donovan and Kernighan book
*/
package main

import (
	"fmt"
)

func main() {
	i := false

	fmt.Printf("initial val: %d", btoi(i))
}

func btoi(b bool) (res int) {
	if b {
		return 1
	} else {
		return 0
	}
}
