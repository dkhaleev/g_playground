/**
temperature converter. Custom and base types
Example from the Donovan and Kernighan book

*/

package tempconv

//FtoC Function converts temperature from Fahrenheit type to Celsius type
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

//CtoF Function converts temperature from Celsius type to Fahrenheit
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
