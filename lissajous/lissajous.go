/**
Generate gif-animated Lissajous figure
Example from the Donovan and Kernighan book
*/

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

const (
	WhiteIndex = 0 //base color
	BlackIndex = 1 //graph color
)

var palette = []color.Color{
	color.RGBA{randomHex(), randomHex(), randomHex(), 0xff},
	color.RGBA{randomHex(), randomHex(), randomHex(), 0xff}}

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // init randomizer
	lissajous(os.Stdout)
}

func randomHex() uint8 {
	rand.Seed(time.Now().UTC().UnixNano()) // init randomizer
	bytes := make([]byte, 1)
	rand.Read(bytes)

	return uint8(rand.Int())
}

func lissajous(out io.Writer) {
	const (
		cycles     = 5     //number of oscillation cycles
		resolution = 0.001 //angle resolution
		size       = 100   //canvas size
		nframes    = 64    //frames in loop
		delay      = 8     //interframe delay 1 = 10ms
	)
	rand.Seed(time.Now().UTC().UnixNano())   // init randomizer
	freq := rand.Float64() * 3               //relative frequency
	animation := gif.GIF{LoopCount: nframes} //init animation struct
	phase := 0.0                             //phases difference

	//iterate frameset
	for i := 0; i < nframes; i++ {

		rectange := image.Rect(0, 0, 2*size+1, 2*size+1) //init frame struct
		img := image.NewPaletted(rectange, palette)      //create frame image

		for j := 0.0; j < cycles*math.Pi*2; j += resolution {
			x := math.Sin(j)
			y := math.Sin(j*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), BlackIndex)
		}

		phase += 0.1
		animation.Delay = append(animation.Delay, delay)
		animation.Image = append(animation.Image, img)
	}

	gif.EncodeAll(out, &animation)
}
