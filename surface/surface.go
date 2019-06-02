/**
Floatin point calculation
Example from the Donovan and Kernighan book
*/

package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320 //canvas size
	cells         = 100      //cells of the surface
	xyrange       = 30.0     //axis range
	// (-xyrange; +xyrange)
	xyscale = width / 2 / xyrange //px per axis
	zscale  = height * 0.4        //px per Z-axis
	angle   = math.Pi / 6         //FOV angle. ~30 degrees
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin 30 deg and cos 30 deg

func main() {
	genSVG()
}

func genSVG() {

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	var err bool
	var ax, ay, bx, by, cx, cy, dx, dy float64

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			if ax, ay, err = corner(i+1, j); !err {
				continue
			}
			if bx, by, err = corner(i, j); !err {
				continue
			}
			if cx, cy, err = corner(i, j+1); !err {
				continue
			}
			if dx, dy, err = corner(i+1, j+1); !err {
				continue
			}

			fmt.Printf("<polygon points='%g, %g, %g, %g, %g, %g, %g, %g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)

		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	//search corner point (x,y) of cell (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	var zero float64
	//calc height of z-surface
	z, ok := f(x, y)
	//isometrically project points on the 2-d canvas(sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	if math.IsInf(sx, 0) || math.IsInf(sy, 0) || math.IsNaN(sx) || math.IsNaN(sy) {
		return zero, zero, false
	}
	return sx, sy, ok
}

func f(x, y float64) (float64, bool) {
	ok := true
	r := math.Hypot(x, y) //distance from (0,0)
	result := math.Sin(r) / r

	//if result is NaN of Inf stop processing
	if math.IsInf(result, 0) || math.IsNaN(result) {
		return result, false
	}

	return result, ok
}
