/*
bare minimal web-server with lissajous output
Example from the Donovan and Kernighan book
*/

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	WhiteIndex = 0
	BlackIndex = 1
)

var palette = []color.Color{
	color.RGBA{randomHex(), randomHex(), randomHex(), 0xff},
	color.RGBA{randomHex(), randomHex(), randomHex(), 0xff}}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UTC().UnixNano()) // init randomizer
	lissajous(w, r)
}

func randomHex() uint8 {
	bytes := make([]byte, 1)
	rand.Read(bytes)

	return uint8(rand.Int())
}

func lissajous(out io.Writer, r *http.Request) {
	var cycles = 5 //number of oscillation cycles
	const (
		resolution = 0.001 //angle resolution
		size       = 100   //canvas size
		nframes    = 64    //frames in loop
		delay      = 8     //interframe delay 1 = 10ms
	)
	//parse form first
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	//iterate form parameters with vals
	for param, val := range r.Form {
		switch param {
		case "cycles":
			// 	fmt.Printf(val[0])
			// 	fmt.Printf("%T\t%s\n", cycles, cycles)
			c, err := strconv.Atoi(val[0])
			if err != nil {
				log.Print(err)
			}
			cycles = c

			break
		}
	}
	// fmt.Print(r.Form["cycles"])
	rand.Seed(time.Now().UTC().UnixNano())   // init randomizer
	freq := rand.Float64() * 3               //relative frequency
	animation := gif.GIF{LoopCount: nframes} //init animation struct
	phase := 0.0                             //phases difference

	//iterate frameset
	for i := 0; i < nframes; i++ {

		rectange := image.Rect(0, 0, 2*size+1, 2*size+1) //init frame struct
		img := image.NewPaletted(rectange, palette)      //create frame image

		for j := 0.0; j < float64(cycles)*math.Pi*2; j += resolution {
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
