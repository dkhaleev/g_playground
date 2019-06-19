/*Complex numbers example.
Chapter 3.3 Example from the Donovan and Kernighan book
*/
package main

import (
	"fmt"
)

func main() {
	var a complex128 = complex(1, 2)
	var b complex128 = complex(3, 4)
	fmt.Println(real(a + b))
	fmt.Println(imag(a + b))
	fmt.Println(a + b)
}
