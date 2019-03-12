/**
temperature converter. Custom and base types
Example from the Donovan and Kernighan book

*/

package conv

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

//FtoC Function converts temperature from Fahrenheit type to Celsius type
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

//CtoF Function converts temperature from Celsius type to Fahrenheit
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g degrees Celsius", c)
}
