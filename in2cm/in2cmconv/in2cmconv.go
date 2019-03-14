/**
Inches to centimeters and vice versa converter. Packages and imports understanding
homework from the Donovan and Kernighan book
*/

package in2cmconv

import (
	"fmt"
)

//custom Inch scale type with base float64
type Inch float64

//custom Centimeter scale type with base float64
type Centimeter float64

//package main constants. Literally just imperial to metric shift
const (
	InchShift Centimeter = 2.54
)

//__toString()
func (c Centimeter) String() string {
	return fmt.Sprintf("%g centimeters", c)
}

//__toString
func (i Inch) String() string {
	return fmt.Sprintf("%g incher", i)
}
