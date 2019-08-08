/*Complex numbers example. Mandelbrot set. Newton method
Chapter 3.3 Example from the Donovan and Kernighan book
*/

package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

const (
	width                  = 4096
	height                 = 4096
	xmin, ymin, xmax, ymax = -2, -2, +2, 2
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < height; i++ {
		y := float64(i)/height*(ymax-ymin) + ymin
		for j := 0; j < width; j++ {
			x := float64(j)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(i, j, newton(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func newton(x complex128) color.Color {
	const (
		iterations = 255
		contrast   = 16
		threshold  = 0.000001
	)

	for n := uint8(0); n < iterations; n++ {
		x = x - (x-1/(cmplx.Pow(x, 3)))/4
		if cmplx.Abs(cmplx.Pow(x, 4)-1) < threshold {
			return map2RGB(255 - contrast*n)
		}
	}

	return color.Black
}

func map2RGB(n uint8) color.Color {

	h := 360 * float64(n) / 256
	x := uint8(255 * (1 - math.Abs(math.Mod(h/60, 2)-1)))

	switch {
	case h < 60:
		return color.RGBA{255, x, 0, 255}
	case h < 120:
		return color.RGBA{x, 255, 0, 255}
	case h < 180:
		return color.RGBA{0, 255, x, 255}
	case h < 240:
		return color.RGBA{0, x, 255, 255}
	case h < 300:
		return color.RGBA{x, 0, 255, 255}
	default:
		return color.RGBA{255, 0, x, 255}
	}

	return color.Black
}
