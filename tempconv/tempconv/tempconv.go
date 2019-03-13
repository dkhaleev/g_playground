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

//Celsius is a custom celsius scale type with base float64
type Celsius float64

//Fahrenheit is a custom fahrenheit scale type with base float64
type Fahrenheit float64

//Kelvin is a custom kelvin scale type with base float64
type Kelvin float64

//Block of some useful celsius scale constants
const (
	AbosluteZeroC Celsius = -273.5
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	KelvinShift   Celsius = 273.5
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g degrees Celsius", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g degrees Fahrenheit", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g degrees Kelvin", k)
}
