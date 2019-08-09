/*strings playground
Chapter 3.5 Example from the Donovan and Kernighan book
*/
package main

import (
	"fmt"
)

func main() {
	s := "This is a test string"

	fmt.Println(len(s))
	fmt.Println(s[1])
	fmt.Println(s[1:3])

	t := s[1:3]
	t += ", All"
	t += " \t tabulated"
	t += " \r\n\v vtabulated"
	t += "\r\n\a sounded"
	fmt.Println(t)

	hexControl := "\xFE"
	fmt.Println(hexControl)

	octControl := "\376"
	fmt.Println(octControl)

	rawLiteral := `This isn't simpe string. It's a [RAW] literal
	Strins formatting and {special} characters % used`
	fmt.Println(rawLiteral)
}
