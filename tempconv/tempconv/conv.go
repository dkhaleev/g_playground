/**
temperature converter. Custom and base types
Example from the Donovan and Kernighan book

*/

package tempconv

//FtoC Function converts temperature from Fahrenheit type to Celsius type
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

//FtoK Function converts temperature from Fahrenheit type to Kelvin type
func FtoK(f Fahrenheit) Kelvin {
	return Kelvin((f - 32 + Fahrenheit(KelvinShift)) * 5 / 9)
}

//CtoF Function converts temperature from Celsius type to Fahrenheit
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

//CtoK Function converts temperature from Celsius type to Kelvin
func CtoK(c Celsius) Kelvin {
	return Kelvin(c + Celsius(KelvinShift))
}

//KtoC Function converts temperature from Kelvin type to Celsius
func KtoC(k Kelvin) Celsius {
	return Celsius(k - Kelvin(KelvinShift))
}

//KtoF Function converts temperature from Kelvin type to Fahrenheit
func KtoF(k Kelvin) Fahrenheit {
	return Fahrenheit(k*5/9 + 32 - Kelvin(KelvinShift))
}
