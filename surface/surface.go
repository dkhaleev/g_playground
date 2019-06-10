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

//polygon structure
type Polygon struct {
	ax, ay float64
	bx, by float64
	cx, cy float64
	dx, dy float64
	az     float64
	bz     float64
	cz     float64
	dz     float64
	zavg   float64
}

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin 30 deg and cos 30 deg
var polygons [cells][cells]Polygon

func main() {
	genSVG()
}

func genSVG() {

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)

	zmin, zmax := genPolygons()

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			red := int(math.Floor(255 * ((polygons[i][j].zavg - zmin) / (zmax - zmin))))
			blue := 255 - red
			green := 0

			fmt.Printf("<polygon points='%g, %g, %g, %g, %g, %g, %g, %g' fill='rgb(%d,%d,%d)'/>\n",
				polygons[i][j].ax, polygons[i][j].ay,
				polygons[i][j].bx, polygons[i][j].by,
				polygons[i][j].cx, polygons[i][j].cy,
				polygons[i][j].dx, polygons[i][j].dy,
				red,
				green,
				blue)

		}
	}
	fmt.Println("</svg>")
}

func genPolygons() (float64, float64) {
	var ok bool
	var zmax, zmin, zavg, az, bz, cz, dz float64

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			if polygons[i][j].ax, polygons[i][j].ay, az, ok = corner(i+1, j); !ok {
				continue
			}
			if polygons[i][j].bx, polygons[i][j].by, bz, ok = corner(i, j); !ok {
				continue
			}
			if polygons[i][j].cx, polygons[i][j].cy, cz, ok = corner(i, j+1); !ok {
				continue
			}
			if polygons[i][j].dx, polygons[i][j].dy, dz, ok = corner(i+1, j+1); !ok {
				continue
			}
			// if ok {
			// 	continue
			// }
			zavg = (az + bz + cz + dz) / 4
			polygons[i][j].zavg = zavg
			polygons[i][j].az = az
			polygons[i][j].bz = bz
			polygons[i][j].cz = cz
			polygons[i][j].dz = dz

			if !math.IsNaN(zavg) {
				zmin = math.Min(zmin, zavg)
				zmax = math.Max(zmax, zavg)
			}
		}
	}

	return zmin, zmax
}

func corner(i, j int) (float64, float64, float64, bool) {
	//search corner point (x,y) of cell (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// var zero float64
	//calc height of z-surface
	z, ok := f(x, y)
	//isometrically project points on the 2-d canvas(sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	if math.IsInf(sx, 0) || math.IsInf(sy, 0) || math.IsNaN(sx) || math.IsNaN(sy) {
		ok = false
	}
	return sx, sy, z, ok
}

func f(x, y float64) (float64, bool) {
	ok := true
	//ex. 3.1
	r := math.Hypot(x, y) //distance from (0,0)
	result := math.Sin(r) / r
	//ex. 3.2
	// result := math.Pow(x, 3) - 3*x*math.Pow(y, 2)

	//if result is NaN of Inf stop processing
	if math.IsInf(result, 0) || math.IsNaN(result) {
		return result, false
	}

	return result, ok
}
