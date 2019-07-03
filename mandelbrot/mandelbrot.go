package main

/*Complex numbers example. Mandelbrot set
Chapter 3.3 Example from the Donovan and Kernighan book
*/

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

const (
	width  = 2048
	height = 2048
)

var imageData [width][height]color.Color
var mask [width / 2][height / 2]color.Color

func main() {
	// fill the image array
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
	)

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin

		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			imageData[px][py] = mandelbrot(z)
		}
	}
	//fill the sampled mask
	for sy := 0; sy < height/2; sy++ {
		for sx := 0; sx < width/2; sx++ {
			//quarterize
			ra, ga, ba, _ := color.Color.RGBA(imageData[sx*2][sy*2])
			rb, gb, bb, _ := color.Color.RGBA(imageData[sx*2+1][sy*2])
			rc, gc, bc, _ := color.Color.RGBA(imageData[sx*2][sy*2+1])
			rd, gd, bd, _ := color.Color.RGBA(imageData[sx*2+1][sy*2+1])
			//equalize
			r := (ra + rb + rc + rd) / 4
			g := (ga + gb + gc + gd) / 4
			b := (ba + bb + bc + bd) / 4
			//fill the mask
			mask[sx][sy] = color.RGBA{uint8(r), uint8(g), uint8(b), 255}
		}
	}
	//render image
	img := image.NewRGBA(image.Rect(0, 0, width/2, height/2))
	for ix := 0; ix < height/2; ix++ {
		for iy := 0; iy < width/2; iy++ {
			img.Set(ix, iy, mask[ix][iy])
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			// return map2RGB{255 - contrast*n}
			return map2RGB(255 - 16*n)
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
