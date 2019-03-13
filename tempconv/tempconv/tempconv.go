/**
main file for types example. Temperature conversion
Example from the Donovan and Kernighan book

*/

/**
temperature converter. Custom and base types
Example from the Donovan and Kernighan book

*/

package tempconv

import "fmt"

//Celsius is a custom celsius type with base float64
type Celsius float64

//Fahrenheit is a custom fahrenheit type with base float64
type Fahrenheit float64

//Block of some useful celsius scale constants
const (
	AbosluteZeroC Celsius = -273.5
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g degrees Celsius", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g degrees Fahrenheit", f)
}
