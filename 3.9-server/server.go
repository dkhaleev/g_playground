/*Complex numbers example. web-server for complex calculations
Chapter 3.8 Example from the Donovan and Kernighan book
*/

package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// var x, y, zoom, visualizer = 0, 0, 1, "mandelbrot"
	var x, y, zoom float64
	var width, height int
	var err error
	var visualizer = "mandelbrot"

	//parse form first
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	//iterate form parameters with vals
	for param, val := range r.Form {
		switch param {
		case "x":
			x, err = strconv.ParseFloat(val[0], 64)
			if err != nil {
				log.Print(err)
				x = 0
			}
			break
		case "y":
			y, err = strconv.ParseFloat(val[0], 64)
			if err != nil {
				log.Print(err)
				y = 0
			}
			break
		case "zoom":
			zoom, err = strconv.ParseFloat(val[0], 64)
			if err != nil {
				log.Print(err)
				zoom = 1
			}
			break
		case "width":
			width, err = strconv.Atoi(val[0])
			if err != nil {
				log.Print(err)
				width = 1024
			}
			break
		case "height":
			height, err = strconv.Atoi(val[0])
			if err != nil {
				log.Print(err)
				height = 1024
			}
			break
		case "visualizer":
			visualizer = val[0]
			break
		}
	}
	render(w, -2/zoom+x, 2/zoom+x, -2/zoom+y, 2/zoom+y, visualizer, width, height)
}

func render(writer http.ResponseWriter, xmin float64, xmax float64, ymin float64, ymax float64, visualizer string, width int, height int) {

	// func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < height; i++ {
		y := float64(i)/float64(height)*(ymax-ymin) + ymin
		for j := 0; j < width; j++ {
			x := float64(j)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			switch visualizer {
			case "newton":
				img.Set(i, j, newton(z))
				break
			case "mandelbrot":
				img.Set(i, j, mandelbrot(z))
				break
			default:
				img.Set(i, j, mandelbrot(z))
				break
			}
		}
	}

	png.Encode(writer, img)
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
