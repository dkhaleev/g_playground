/*
Fahrenheit to Celcius converter
Example from the Donovan and Kernighan book
*/

package main

import (
	"fmt"
)

func main() {
	const freezingF, boilingF, burningF = 32.0, 212.0, 451.0
	fmt.Printf("Boiling temperature of water %g Fahrenheit and %g Celsius \n", boilingF, fToC(boilingF))
	fmt.Printf("Freezing temperature of water %g Fahreheit and %g Celsius \n", freezingF, fToC(freezingF))
	fmt.Printf("Burning temerature of paper %g Fahrenheit and %g Celsius \n", burningF, fToC(burningF))
}

//function converts fahrenheit degrees to celsius
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
