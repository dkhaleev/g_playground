/**
Inches to centimeters and vice versa converter. Packages and imports understanding
homework from the Donovan and Kernighan book
*/

package in2cmconv

//In2cm Function converts inches from imperial to centimeter metric
func In2cm(i Inch) Centimeter {
	return Centimeter(i * Inch(InchShift))
}

//Cm2in function converts centimeter from metric to inches imperial
func Cm2in(c Centimeter) Inch {
	return Inch(c / InchShift)
}
