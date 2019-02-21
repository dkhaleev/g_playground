/*
Fahrenheit to Celcius converter
Example from the Donovan and Kernighan book
*/

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	const freezingF, boilingF, burningF = 32.0, 212.0, 451.0

	if len(os.Args) < 3 {
		fmt.Println("Too few arguments. Temp or Scale missing")
		os.Exit(1)
	}

	scale := os.Args[2]

	temp, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Print(err)
	}

	switch scale {
	case "F":
		fmt.Println("FtC conversion")
		fmt.Printf("%g degrees Fahrenheit equals %g Celsius\n", temp, ftoC(temp))
		break
	case "C":
		fmt.Println("CtF conversion")
		fmt.Printf("%g degrees Celsius equals %g Fahrenheit\n", temp, ctoF(temp))
		break
	default:
		fmt.Println("Unknown scale")
		os.Exit(1)
	}
}

//function converts fahrenheit degrees to celsius
func ftoC(f float64) float64 {
	return (f - 32) * 5 / 9
}

//function converts celsuis to fahrenheit
func ctoF(c float64) float64 {
	return (c * 9 / 5) + 32
}
